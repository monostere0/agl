package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func (AGL) getLogEvents(searchTerm string) []types.OutputLogEvent {
	var r []types.OutputLogEvent
	var cl cloudwatchlogs.Client

	initClient(&cl)

	lgs := getLogStreams(searchTerm, &cl)

	// make a buffered channel for log streams requests
	c := make(chan []types.OutputLogEvent, len(lgs))

	// iterate through the log streams and
	// launch a goroutine for getting its events and store them in channel c
	for _, lv := range lgs {
		for _, sv := range lv.logStreams {
			go func(lv LogStreamsOutput, sv types.LogStream) {
				output, err := cl.GetLogEvents(context.TODO(), &cloudwatchlogs.GetLogEventsInput{
					LogStreamName: sv.LogStreamName,
					LogGroupName:  &lv.logGroupName,
				})

				if err != nil {
					log.Fatal(err)
				}

				c <- output.Events
			}(lv, sv)

		}

	}

	// write all channel results into the return slice
	for i := 0; i < len(lgs); i++ {
		r = append(r, <-c...)
	}

	return r
}
