package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/tebeka/nrsc"
)

const (
	address = "127.0.0.1:8081"
)

var (
	messages chan botResponse
)

func init() {
	banner := fmt.Sprintf("\n\t.:: Please go to http://%s/index.html ::.\n", address)
	fmt.Println(banner)
}

func main() {
	messages = make(chan botResponse, 10)

	http.HandleFunc("/webhook", forwardMessages)
	http.HandleFunc("/bottoken/sendMessage", mockTelegram)
	nrsc.Handle("/")

	handler := accessLog(http.DefaultServeMux)

	if err := http.ListenAndServe(address, handler); err != nil {
		fmt.Println(err)
	}
}
