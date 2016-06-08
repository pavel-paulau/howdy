package main

import (
	"fmt"
	"net/http"
)

const (
	address = "127.0.0.1:8080"
	greeting = `
This Telegram chat bot was created for demonstration purposes only.

It receives incoming updates via a webhook with the following address:

	http://127.0.0.1:8080/webhook
`
)

func init() {
	fmt.Println(greeting)
}

func main() {
	http.HandleFunc("/webhook", chatBot)

	if err := http.ListenAndServe(address, http.DefaultServeMux); err != nil {
		fmt.Println(err)
	}
}
