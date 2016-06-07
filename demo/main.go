package main

import (
	"fmt"
	"net/http"
)

const (
	address = "127.0.0.1:8080"
)

func main() {
	http.HandleFunc("/webhook", chatBot)

	if err := http.ListenAndServe(address, http.DefaultServeMux); err != nil {
		fmt.Println(err)
	}
}
