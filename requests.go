package main

import (
	"io/ioutil"
	"net/http"
)

/*
getRequest is a convenience function to make GET requests
*/
func getRequest(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	handleErrorAndPanic(err)

	client := http.Client{}

	resp, err := client.Do(req)
	handleErrorAndPanic(err)

	body, err := ioutil.ReadAll(resp.Body)
	handleErrorAndPanic(err)

	err = resp.Body.Close()
	handleErrorAndPanic(err)

	return string(body)
}
