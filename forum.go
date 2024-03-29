package main

import (
	"encoding/json"
	"fmt"
)

/*
ClassForum is a struct component to store a list of unique commenters based on given forum name. It scrolls through infinite
scroll system of disqus and acquires every commenter's username, profile URL, private or public profile, etc. meta info.
*/
type ClassForum struct {
	Users     []User
	ForumName string
	Key       string
	IsPrivate bool
}

/*
getCommenters parses the json response from the infinite scroll mechanics on disqus and stores all found users.
Note: By default we skip private or blocked profiles, as well as those that have missing profile links in the json
response.
TODO: Try to eliminate duplicates if any? Probably not necessary.
*/
func (m *ClassForum) getCommenters() {
	debugLog("func ClassForum.getCommenters() called with params [", m.Key, "]")

	cursor := ""

hereWeGoAgain:
	var w UserListResponseModel

	err := json.Unmarshal([]byte(getRequest(fmt.Sprintf(usersByForumURL, m.ForumName, cursor, m.Key))), &w)
	handleErrorAndPanic(err)

	for _, v := range w.Response {
		var u User

		// Skip if any of following conditions
		if v.ProfileURL == "" || v.IsPrivate || v.IsBlocked {
			continue
		}

		u.profileURL = v.ProfileURL
		u.location = v.Location
		u.joinedAt = v.JoinedAt
		u.reputation = v.Reputation
		u.username = v.Username

		m.Users = append(m.Users, u)
	}

	if w.Cursor.HasNext {
		cursor = w.Cursor.Next
		debugLog("Getting next batch of users for", m.ForumName, "next is", w.Cursor.Next)
		goto hereWeGoAgain
	}

	debugLog("func ClassForum.getCommenters() returning with users set [", len(m.Users), m.Users, "]")
}

/*
isPrivate verifies whether given forum name is private
*/
func (m *ClassForum) isPrivate() {
	debugLog("func ClassForum.isPrivate() called")

	var r ForumActivityResponseModel

	err := json.Unmarshal([]byte(getRequest(fmt.Sprintf(linksByDirectThreadURL, m.ForumName, "", m.Key))), &r)
	handleErrorAndPanic(err)

	if len(r.Response.Activities) == 0 {
		m.IsPrivate = true
	} else {
		m.IsPrivate = false
	}

	debugLog("func ClassForum.isPrivate() returning with m.IsPrivate set to", m.IsPrivate, len(r.Response.Activities))
}
