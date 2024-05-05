package main

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type AGL struct{}

type LogStreamsOutput struct {
	logGroupName string
	logStreams   []types.LogStream
}
