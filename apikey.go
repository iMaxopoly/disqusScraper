package main

import "fmt"

/*
ClassAPI is required per request. This is likely IP-restricted, making me wonder if this may have Rate-limiting as in the
official API docs and may have additional limitations. We can possibly avoid such restrictions using proxies.
TODO: Find out if requests have rate limiting and/or other restrictions.
*/
type ClassAPI struct {
	Key string
}

/*
solveInitializerJS locates hint to receive the full API key in solveMainJS
*/
func (m *ClassAPI) solveInitializerJS() string {
	debugLog("func ClassAPI.solveInitializerJS() called")
	serverResponse := getRequest(mAPIKeyInitializerJsURL)

	result := patternMatch(serverResponse, patternInitializerJS, '/')

	debugLog("func ClassAPI.solveInitializerJS() returning with result [", result, "]")
	return result
}

/*
solveMainJS locates and sets the full API key we need to begin making requests
*/
func (m *ClassAPI) solveMainJS(hint *string) {
	debugLog("func ClassAPI.solveMainJS() called with params [", *hint, "]")
	serverResponse := getRequest(fmt.Sprintf("https://a.disquscdn.com/next/%s/home/js/main.js", *hint))

	m.Key = patternMatch(serverResponse, patternMainJS, '"')

	debugLog("func ClassAPI.solveMainJS() sets APIKey.Value and returns  [", m.Key, "]")
}

/*
setup is a macro that runs all the steps to fetch the APIKey Value
*/
func (m *ClassAPI) setup() {
	debugLog("func ClassAPI.setup() called")

	apiHint := m.solveInitializerJS()
	m.solveMainJS(&apiHint)

	debugLog("func ClassAPI.setup() returning")
}
