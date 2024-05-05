package main

import (
	"fmt"
	"os"
)

func main() {
	// instead of setting in the terminal
	os.Setenv("AWS_REGION", "eu-central-1")

	agl := AGL{}
	logStreams := agl.getLogStreams()

	for _, v := range logStreams {
		fmt.Println(v.logGroupName)
	}
}
