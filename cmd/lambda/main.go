package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mctofu/pushgate"
)

func main() {
	key := os.Getenv("PUSHOVER_KEY")
	rcpt := os.Getenv("PUSHOVER_RCPT")
	h := &handler{
		sender:    pushgate.NewPushoverSender(key),
		recipient: rcpt,
	}
	lambda.Start(h.handle)
}

type handler struct {
	sender    *pushgate.PushoverSender
	recipient string
}

func (h *handler) handle(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		r := record.SNS
		if err := h.sender.Send(h.recipient, &pushgate.Message{Title: r.Subject, Body: r.Message}); err != nil {
			if err, ok := err.(*pushgate.RetryableError); ok {
				return err
			}
			// we are doing something wrong, try to notify/log but don't return an error so SNS doesn't retry
			log.Printf("Failed to send push: %v\n", err)
			if err := h.sender.Send(h.recipient, &pushgate.Message{Title: "PushGate error", Body: err.Error()}); err != nil {
				log.Printf("Failed to send error push: %v\n", err)
			}
		}
	}

	return nil
}
