package main

const (
	/*
		mAPIKeyInitializerJsURL is the URL location for initializer.js which sets up the API key for our project scope.
	*/
	mAPIKeyInitializerJsURL = "https://a.disquscdn.com/next/current/home/js/apps/initializer.js"

	/*
		patternInitializerJS is a helper pattern to locate hint to fetch the public API key for given IP
	*/
	patternInitializerJS = "baseUrl:\"//a.disquscdn.com/next/"

	/*
		whatWeCanDMCAURL is a structure URL that *SHOULD* point to the URL of the thread where comments are posted and are
		therefore able to be used for DMCA reasons
	*/
	whatWeCanDMCAURL = "https://disqus.com/home/discussion/%s/%s"

	/*
		patternMainJS is a helper pattern to locate API Key for given IP
	*/
	patternMainJS = "keys:{api:\""

	/*
		patternFeed is a helper pattern to that locates links for given json response in a Comment Activity API reply
	*/
	patternFeed = "\"feed\":\""
)

var (
	/*
	   configDebug switches between the program being verbose or silent
	*/
	configDebug = true
	/*
		usersByForumURL is the public facing API URL that provides us with a json response containing maximum 100 users
		at a time for given forum name
	*/
	usersByForumURL = "https://disqus.com/api/3.0/forums/listMostLikedUsers?forum=%s&limit=100&cursor=%s&api_key=%s"

	/*
		linksByCommentActivityURL is the public facing API URL that provides us with a json response containing
	*/
	linksByCommentActivityURL = "https://disqus.com/api/3.0/timelines/activities?type=profile&index=comments&target=user%%3Ausername%%3A%s&cursor=%s&limit=100&api_key=%s"

	/*
		linksByDirectThreadURL is the public facing API URL that provides us with complete thread listing of the given
		forum considering it isn't private
	*/
	linksByDirectThreadURL = "https://disqus.com/api/3.0/timelines/activities?type=profile&target=forum%%3A%s&cursor=%s&limit=100&api_key=%s"
)
