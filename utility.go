package main

import (
	"log"
	"os"
	"strings"
	"sync"
)

/*
handleErrorAndPanic is a convenience error handler that panics
*/
func handleErrorAndPanic(err error, msg ...string) {
	if err != nil {
		log.Panicln(err, msg)
	}
}

/*
debugLog prints given messages if configDebug const is true
*/
func debugLog(msg ...interface{}) {
	if configDebug {
		log.Println(msg)
	}
}

/*
patternMatch is convenience function to match pattern and extract what is needed from provided data
*/
func patternMatch(serverResponse, pattern string, dealBreaker byte) string {
	debugLog("func patternMatch() called with params [", serverResponse[:20], pattern, string(dealBreaker), "]")

	var whatWeAreLookingFor string

	mIndex := strings.Index(serverResponse, pattern)

	for i := mIndex + len(pattern); i < len(serverResponse); i++ {
		if serverResponse[i] == dealBreaker {
			whatWeAreLookingFor = serverResponse[mIndex+len(pattern) : i]
			break
		}
	}

	debugLog("func patternMatch() returning with result [", whatWeAreLookingFor, "]")
	return whatWeAreLookingFor
}

var writeLock sync.Mutex

func writeToFile(filename, content string) {
	writeLock.Lock()
	defer writeLock.Unlock()

	debugLog("Writing", filename, content)
	if strings.TrimSpace(filename) == "" || strings.TrimSpace(content) == "" {
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	handleErrorAndPanic(err)

	_, err = file.WriteString(content + "\n")
	handleErrorAndPanic(err)

	err = file.Close()
	handleErrorAndPanic(err)
}
