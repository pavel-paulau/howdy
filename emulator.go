package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	log "gopkg.in/inconshreveable/log15.v2"
)

type ChatMessage struct {
	Text      string `json:"text"`
	FirstName string `json:"firstName"`
	UserName  string `json:"userName"`
	UserId    int    `json:"userId"`
	Phone     string `json:"phone"`
	Webhook   string `json:"webhook"`
}

func readBody(req *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(payload)
	if err != nil {
		log.Error("failed to parse body", "err", err)
	}
	return err
}

func forwardMessages(rw http.ResponseWriter, req *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Error("failed to upgrade to websockets", "err", err)
		return
	}

	go func() {
		defer conn.Close()

		for {
			botResponse := <-messages

			if err := conn.WriteJSON(&botResponse); err != nil {
				log.Error("failed to write data", "err", err)
				break
			}
		}
	}()

	go func() {
		defer conn.Close()

		for {
			var message ChatMessage
			if err := conn.ReadJSON(&message); err != nil {
				log.Error("failed to read data", "err", err)
				break
			} else {
				sendUpdateToBot(message)
			}
		}
	}()
}

func sendUpdateToBot(message ChatMessage) {
	update := Update{
		Message: Message{
			Text: message.Text,
			From: User{
				FirstName: message.FirstName,
				ID:        message.UserId,
				Username:  message.UserName,
			},
		},
		UpdateID: 0,
	}

	if _, err := sendJSON(message.Webhook, &update); err != nil {
		log.Error("failed to send update", "err", err)
	}
}

func mockTelegram(rw http.ResponseWriter, req *http.Request) {
	var botResponse BotResponse
	readBody(req, &botResponse)

	messages <- botResponse

	response := TelegramResponse{
		OK:     true,
		Result: Message{}, // FIXME
	}
	json.NewEncoder(rw).Encode(&response)
}
