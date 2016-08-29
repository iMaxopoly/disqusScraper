package main

/*
User is a simple model to store user information if required in further extensions of the project
Note: More meta data can be obtained, the model provided as-is, is simply what was deemed important at the time of
programming.
*/
type User struct {
	username   string
	profileURL string
	location   string
	joinedAt   string
	reputation float64
}
