package main

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func getLogStreams(searchTerm string, cl *cloudwatchlogs.Client) []LogStreamsOutput {
	var r []LogStreamsOutput

	output, err := cl.DescribeLogGroups(context.TODO(), &cloudwatchlogs.DescribeLogGroupsInput{})
	lgs := filterLogGroupsByName(searchTerm, output.LogGroups)

	if len(lgs) == 0 {
		panic("No logs were found with the name " + searchTerm)
	}

	// make a buffered channel for log streams requests
	c := make(chan LogStreamsOutput, len(output.LogGroups))

	if err != nil {
		log.Fatal(err)
	}

	// iterate through the log groups and
	// launch a goroutine for getting its logstreams and store them in channel c
	for _, v := range lgs {
		go func(lg string) {
			output, err := cl.DescribeLogStreams(context.TODO(), &cloudwatchlogs.DescribeLogStreamsInput{
				LogGroupName: &lg,
			})

			if err != nil {
				log.Fatal(err)
			}

			c <- LogStreamsOutput{
				logGroupName: lg,
				logStreams:   output.LogStreams,
			}

		}(string(*v.LogGroupName))
	}

	// write all channel results into the return slice
	for i := 0; i < len(lgs); i++ {
		r = append(r, <-c)
	}

	return r
}

func filterLogGroupsByName(searchTerm string, logGroups []types.LogGroup) []types.LogGroup {
	var r []types.LogGroup
	for _, v := range logGroups {
		if strings.Contains(string(*v.LogGroupName), searchTerm) {
			r = append(r, v)
		}
	}

	return r
}
