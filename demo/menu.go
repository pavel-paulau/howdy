package main

import (
	"fmt"

	"github.com/pavel-paulau/howdy/telegram"
)

type menuItem struct {
	Emoji string
	Text  string
}

func (i *menuItem) Label() string {
	if i.Emoji != "" {
		return fmt.Sprintf("%s %s", i.Emoji, i.Text)
	}
	return i.Text
}

var (
	mainMenu = []menuItem{
		{Emoji: "\xF0\x9F\x9A\x91", Text: "Ambulance service"},
		{Emoji: "\xF0\x9F\x9A\x92", Text: "Fire service"},
		{Emoji: "\xF0\x9F\x9A\x93", Text: "Police service"},
	}

	phone = menuItem{Emoji: "\xF0\x9F\x93\xB1", Text: "Send my phone number"}
)

func newKeyboard(items []menuItem) telegram.ReplyKeyboardMarkup {
	buttons := [][]telegram.KeyboardButton{
		{},
	}

	for _, item := range items {
		button := telegram.KeyboardButton{Text: item.Label()}
		buttons[0] = append(buttons[0], button)
	}

	return telegram.ReplyKeyboardMarkup{Keyboard: buttons}
}

func contactRequest() telegram.ReplyKeyboardMarkup {
	buttons := [][]telegram.KeyboardButton{
		{
			{phone.Label(), true},
		},
	}

	return telegram.ReplyKeyboardMarkup{Keyboard: buttons}
}
