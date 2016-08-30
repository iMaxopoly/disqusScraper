package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/go-playground/pool.v3"
)

/*
ClassLinkCollector is our component that automatically allocates workers to range of users, pulling links from where
they have posted based on given forum name(or everything) and ultimately filters for duplicates.
*/
type ClassLinkCollector struct {
	Workers     uint
	ForumName   string
	ResultCount uint
	Key         string
	NoNoise     bool
	Users       []User
}

/*
processPublic fetches all the links from all the public forum of a given forum name.
*/
func (m *ClassLinkCollector) processPublic() {
	debugLog("func ClassLinkCollector.processPublic() called")
	cursor := ""

	var links []string

hereWeGoAgain:
	var c Cursor

	body := getRequest(fmt.Sprintf(linksByDirectThreadURL, m.ForumName, cursor, m.Key))

	err := json.Unmarshal([]byte(body), &c)
	handleErrorAndPanic(err)

notAllOverAgain:
	slugIndex := strings.Index(body, patternFeed)
	if slugIndex != -1 {
		// We get Feed URL, which we are not interested in
		feed := patternMatch(body, patternFeed, '"')

		// <-Surgery Begins-> This step entirely exists because of paranoia
		feed = strings.TrimPrefix(feed, "https://")

		var forumName string
		for i := 0; i < len(feed); i++ {
			if feed[i] == '.' {
				forumName = feed[:i]
				feed = strings.TrimSuffix(strings.TrimPrefix(feed[i:], ".disqus.com/"), "/latest.rss")
				break
			}
		}
		// <-Surgery Ends->

		links = append(links, fmt.Sprintf(whatWeCanDMCAURL, forumName, feed))

		body = body[slugIndex+len(patternFeed)-1:]
		goto notAllOverAgain
	}

	if c.HasNext {
		cursor = c.Next
		debugLog("Getting next batch of links for", m.ForumName, "next is", c.Next)
		goto hereWeGoAgain
	}

	var countLinks uint
	// Write to File
	for i := 0; i < len(links); i++ {
		writeToFile(m.ForumName+".txt", links[i])
		countLinks++
	}

	m.ResultCount = countLinks
	debugLog("func ClassLinkCollector.processPublic() returning with", countLinks, "acquired")
}

/*
processPrivate spawns multiple go-routines to fetch all the links from all the users of a given forum name.
It spawns go-routines based on the workers provided.
*/
func (m *ClassLinkCollector) processPrivate() {
	debugLog("func ClassLinkCollector.processPrivate() called")
	p := pool.NewLimited(m.Workers)
	defer p.Close()

	batch := p.Batch()

	go func() {
		for i := 0; i < len(m.Users); i++ {
			batch.Queue(m.collect(m.Users[i].username))
		}

		batch.QueueComplete()
	}()

	var countLinks uint
	for links := range batch.Results() {

		if err := links.Error(); err != nil {
			batch.Cancel()
			handleErrorAndPanic(err)
		}

		// Write to File
		for i := 0; i < len(links.Value().([]string)); i++ {
			writeToFile(m.ForumName+".txt", links.Value().([]string)[i])
			countLinks++
		}
	}

	m.ResultCount = countLinks
	debugLog("func ClassLinkCollector.processPrivate() returning with", countLinks, "acquired")
}

/*
collect is can be considered as a job function that each go-routine must complete spawned from process
*/
func (m *ClassLinkCollector) collect(username string) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		debugLog("func ClassLinkCollector.collect() called with params [", username, "]")
		if wu.IsCancelled() {
			// return values not used
			return nil, nil
		}

		cursor := ""

		var links []string

	hereWeGoAgain:
		var c Cursor

		body := getRequest(fmt.Sprintf(linksByCommentActivityURL, username, cursor, m.Key))

		err := json.Unmarshal([]byte(body), &c)
		if err != nil {
			return []string{}, err
		}

	notAllOverAgain:
		slugIndex := strings.Index(body, patternFeed)
		if slugIndex != -1 {
			// We get Feed URL, which we are not interested in
			feed := patternMatch(body, patternFeed, '"')

			// <-Surgery Begins-> This step entirely exists because of paranoia
			feed = strings.TrimPrefix(feed, "https://")

			var forumName string
			for i := 0; i < len(feed); i++ {
				if feed[i] == '.' {
					forumName = feed[:i]
					feed = strings.TrimSuffix(strings.TrimPrefix(feed[i:], ".disqus.com/"), "/latest.rss")
					break
				}
			}
			// <-Surgery Ends->

			// Depending on -nonoise switch, if true, only picks up links that belong to given forum name, excludes
			// everything else
			if m.NoNoise {
				if forumName == m.ForumName {
					links = append(links, fmt.Sprintf(whatWeCanDMCAURL, forumName, feed))
				}
			} else {
				links = append(links, fmt.Sprintf(whatWeCanDMCAURL, forumName, feed))
			}

			body = body[slugIndex+len(patternFeed)-1:]
			goto notAllOverAgain
		}

		if c.HasNext {
			cursor = c.Next
			debugLog("Getting next batch of links for", username, "next is", c.Next)
			goto hereWeGoAgain
		}

		debugLog("func ClassLinkCollector.collect() called with param", username, "returning with [", len(links), "] acquired")
		return links, nil // everything ok, send nil, error if not
	}
}
