package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type LogStreamsOutput struct {
	logGroupName string
	logStreams   []types.LogStream
}

type AGL struct{}

func (AGL) init() cloudwatchlogs.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("daniel-dev-gigs"))

	if err != nil {
		log.Fatal(err)
	}

	return *cloudwatchlogs.NewFromConfig(cfg)
}

func (a AGL) getLogStreams() []LogStreamsOutput {
	// get the log group names
	var r []LogStreamsOutput

	client := a.init()
	output, err := client.DescribeLogGroups(context.TODO(), &cloudwatchlogs.DescribeLogGroupsInput{})
	// make a buffered channel for log streams requests
	c := make(chan LogStreamsOutput, len(output.LogGroups))

	if err != nil {
		log.Fatal(err)
	}

	// iterate through the log groups and
	// launch a goroutine for getting its logstreams and store them in channel c
	for _, v := range output.LogGroups {
		go func(lg string) {
			output, err := client.DescribeLogStreams(context.TODO(), &cloudwatchlogs.DescribeLogStreamsInput{
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

	for i := 0; i < len(output.LogGroups); i++ {
		r = append(r, <-c)
	}

	return r
}
