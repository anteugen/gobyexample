package main

import (
	"fmt"
)

func selector(messages chan string) {
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No activity")
	}
}

func main() {
	messages := make(chan string, 1)

	selector(messages)

	messages <- "Message 1"

	selector(messages)

	selector(messages)
}