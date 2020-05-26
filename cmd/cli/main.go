package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mctofu/pushgate"
)

func main() {
	rcpt := flag.String("recipient", "", "pushover recipient key (required)")
	title := flag.String("title", "", "message title")
	body := flag.String("body", "", "message body (required)")
	flag.Parse()

	if *rcpt == "" || *body == "" {
		flag.Usage()
		os.Exit(1)
	}

	key := os.Getenv("PUSHOVER_KEY")
	sender := pushgate.NewPushoverSender(key)

	if err := sender.Send(*rcpt, &pushgate.Message{Title: *title, Body: *body}); err != nil {
		fmt.Printf("Failed to send: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Push sent!\n")
}
