package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	lambda.Start(handler)
}

func handler(_ context.Context, kinesisEvent events.KinesisEvent) error {
	for _, record := range kinesisEvent.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)

		log.Printf("%s Data = %s \n", record.EventName, dataText)
		if dataText == "InvalidData" {
			log.Printf("%s Invalid!!! Data = %s \n", record.EventName, dataText)
			return errors.New("Invalid!!!")
		}
	}
	return nil
}
