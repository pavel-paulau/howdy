package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/tebeka/nrsc"
)

const (
	address = ":8081"
	banner  = "\n\t.:: Please go to http://127.0.0.1:8081/index.html ::.\n"
)

var (
	messages chan botResponse
)

func main() {
	messages = make(chan botResponse, 10)

	http.HandleFunc("/webhook", forwardMessages)
	http.HandleFunc("/bottoken/sendMessage", mockTelegram)
	nrsc.Handle("/")

	handler := accessLog(http.DefaultServeMux)

	fmt.Println(banner)
	if err := http.ListenAndServe(address, handler); err != nil {
		fmt.Println(err)
	}
}
