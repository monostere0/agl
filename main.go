package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	c := *cloudwatchlogs.NewFromConfig(cfg)

	agl := AGL{client: c}
	logStreams := agl.getLogEvents(os.Args[1])

	for _, v := range logStreams {
		fmt.Println(string(*v.Message))
	}
}
