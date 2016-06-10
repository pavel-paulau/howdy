package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/pavel-paulau/howdy/telegram"
)

var (
	updateID int
)

func init() {
	messages = make(chan botResponse, 10)
}

type chatMessage struct {
	Text      string `json:"text"`
	FirstName string `json:"firstName"`
	UserName  string `json:"userName"`
	UserID    int    `json:"userId"`
	Phone     string `json:"phone"`
	Webhook   string `json:"webhook"`
}

type telegramResponse struct {
	OK bool `json:"ok"`
}

type botResponse struct {
	ChatID      int         `json:"chat_id"`      // Unique identifier for the target chat
	Text        string      `json:"text"`         // Text of the message to be sent
	ReplyMarkup interface{} `json:"reply_markup"` // Additional interface options.
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

	isOpen := true
	var myUser int

	go func() {
		defer func() {
			isOpen = false
			conn.Close()
		}()

		for {
			var message chatMessage
			if err := conn.ReadJSON(&message); err != nil {
				log.Error("failed to read data", "err", err)
				break
			}
			myUser = message.UserID
			sendUpdateToBot(message)
		}
	}()

	go func() {
		defer conn.Close()

		for {
			response := <-messages

			if !isOpen {
				messages <- response
				break
			}

			if response.ChatID != myUser {
				messages <- response
				continue
			}

			if err := conn.WriteJSON(&response); err != nil {
				log.Error("failed to write data", "err", err)
				break
			}
		}
	}()
}

func sendUpdateToBot(message chatMessage) {
	updateID++

	update := telegram.Update{
		Message: telegram.Message{
			Text: message.Text,
			From: telegram.User{
				FirstName: message.FirstName,
				ID:        message.UserID,
			},
			Contact: telegram.Contact{
				FirstName:   message.FirstName,
				PhoneNumber: message.Phone,
			},
		},
		UpdateID: updateID,
	}

	if _, err := SendJSON(message.Webhook, &update); err != nil {
		log.Error("failed to send update", "err", err)
	}
}

func mockTelegram(rw http.ResponseWriter, req *http.Request) {
	var br botResponse
	var tr telegramResponse

	if err := ReadJSON(req, &br); err != nil {
		log.Error(err.Error())
		tr = telegramResponse{OK: false}
	} else {
		messages <- br
		tr = telegramResponse{OK: true}
	}

	json.NewEncoder(rw).Encode(&tr)
}
