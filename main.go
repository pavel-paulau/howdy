package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	address = "127.0.0.1:8081"
)

var (
	messages chan BotResponse
)

func init() {
	banner := fmt.Sprintf("\n\t.:: Serving http://%s/ ::.\n", address)
	fmt.Println(banner)
}

func main() {
	messages = make(chan BotResponse, 10)

	r := mux.NewRouter()
	r.HandleFunc("/webhook", forwardMessages)
	r.HandleFunc("/bottoken/sendMessage", mockTelegram).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./app")))

	http.Handle("/", r)
	http.ListenAndServe(address, accessLog(http.DefaultServeMux))
}
