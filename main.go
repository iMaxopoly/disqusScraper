package main

/***
 *          _ _                      _____
 *         | (_)                    / ____|
 *       __| |_ ___  __ _ _   _ ___| (___   ___ _ __ __ _ _ __   ___ _ __
 *      / _` | / __|/ _` | | | / __|\___ \ / __| '__/ _` | '_ \ / _ \ '__|
 *     | (_| | \__ \ (_| | |_| \__ \____) | (__| | | (_| | |_) |  __/ |
 *      \__,_|_|___/\__, |\__,_|___/_____/ \___|_|  \__,_| .__/ \___|_|
 *                     | |                               | |
 *                     |_|                               |_|
 *
 *  Contact:
 *  Manish Prakash Singh
 *  contact@kryptodev.com
 *  Skype: kryptodev
 *
 *  License: This work is licensed under a Creative Commons Attribution-ShareAlike 4.0 International License.
 */

import (
	"flag"
	"fmt"
)

func main() {

	// Command-line setup
	forumName := flag.String("forum", "fiestaonline", "Forum Name; Defaults to \"fiestaonline\" eg. kissanime from https://disqus.com/home/forum/kissanime/commenters/")
	parseUsers := flag.Bool("parseusers", false, "Decides whether to get forum commenters, and then get all threads they have posted on; Defaults to false")
	noNoise := flag.Bool("nonoise", true, "Decides whether to only pick up links that stem from given forum name, or ignore that and pick up everything; Defaults to true")
	workers := flag.Uint("workers", 10, "Number of workers making requests simultaneously and getting the links; Defaults to 10")
	debug := flag.Bool("debug", true, "Whether you want verbosity about what the program is doing; Defaults to true")
	flag.Parse()

	// To be loud or to not be loud!
	configDebug = *debug

	var api ClassAPI
	api.setup() // We setup API to make requests

	var f ClassForum
	f.Key = api.Key
	f.ForumName = *forumName
	f.isPrivate() // We check if forum is private

	if f.IsPrivate {
		// We acquire complete list of commenters
		f.getCommenters()
	}

	// We begin spawning our workers to get links and saving them in a text file named -your forum name-.txt
	var mrLinkCollector ClassLinkCollector
	mrLinkCollector.Key = api.Key
	mrLinkCollector.ForumName = f.ForumName

	if f.IsPrivate || *parseUsers {
		mrLinkCollector.Workers = *workers
		mrLinkCollector.Users = f.Users
		mrLinkCollector.NoNoise = *noNoise
		mrLinkCollector.processPrivate()
	} else {
		mrLinkCollector.processPublic()
	}

	debugLog(fmt.Sprintf("[[%s WORK COMPLETE]]", *forumName))
}
