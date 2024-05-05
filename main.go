package main

import (
	"fmt"
	"os"
)

func main() {
	// instead of setting in the terminal
	os.Setenv("AWS_REGION", "eu-central-1")
	os.Setenv("AWS_PROFILE", "daniel-dev-gigs")

	agl := AGL{}
	logStreams := agl.getLogStreams("createArticle")

	for _, v := range logStreams {
		fmt.Println(v.logGroupName)
	}
}
