package main

import (
	"fmt"
	"net/http"

	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/pavel-paulau/howdy/telegram"
)

const (
	sendMessageURL = "http://127.0.0.1:8081/bottoken/sendMessage"
)

func chatBot(rw http.ResponseWriter, req *http.Request) {
	var update telegram.Update
	if err := ReadJSON(req, &update); err != nil {
		return
	}

	var chatID = update.Message.From.ID

	log.Info("received message", "from", chatID, "text", update.Message.Text)

	switch update.Message.Text {
	case mainMenu[0].Label(), mainMenu[1].Label(), mainMenu[2].Label():
		sendMessage(chatID, mConfirm)
		sendMessage(chatID, mContactNeed)
		kb := contactRequest()
		sendKeyboard(chatID, mRequestContact, kb)
	case phone.Label():
		msg := fmt.Sprintf(mGotNumber, update.Message.Contact.PhoneNumber)
		sendMessage(chatID, msg)
		sendMessage(chatID, mGoodbye)
	default:
		msg := fmt.Sprintf(mHi, update.Message.From.FirstName)
		sendMessage(chatID, msg)
		kb := newKeyboard(mainMenu)
		sendKeyboard(chatID, mSelect, kb)
	}
}

type chatMessage struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func sendMessage(chatID int, text string) {
	log.Info("Sending message", "to", chatID, "text", text)

	message := chatMessage{chatID, text, "HTML"}

	if err := SendJSON(sendMessageURL, &message); err != nil {
		log.Error("failed to send update", "err", err)
	}
}

type chatMessageWithKeyboard struct {
	ChatID        int                          `json:"chat_id"`
	Text          string                       `json:"text"`
	ReplyKeyboard telegram.ReplyKeyboardMarkup `json:"reply_markup"`
}

func sendKeyboard(chatID int, text string, keyboard telegram.ReplyKeyboardMarkup) {
	log.Info("Sending message", "to", chatID, "text", text)

	message := chatMessageWithKeyboard{chatID, text, keyboard}

	if err := SendJSON(sendMessageURL, &message); err != nil {
		log.Error("failed to send update", "err", err)
	}
}
