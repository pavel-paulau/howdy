package main

type User struct {
	FirstName string `json:"first_name"`
	ID        int    `json:"id"`
	Username  string `json:"username"`
}

type Chat struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type Contact struct {
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
}

type Message struct {
	Chat      Chat    `json:"chat"`
	Contact   Contact `json:"contact"`
	Date      int     `json:"date"`
	From      User    `json:"from"`
	MessageID int     `json:"message_id"`
	Text      string  `json:"text"`
}

type Update struct {
	Message  Message `json:"message"`
	UpdateID int     `json:"update_id"`
}

type TelegramResponse struct {
	OK     bool    `json:"ok"`
	Result Message `json:"result"`
}

type BotResponse struct {
	ChatID      int         `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type ReplyMarkup struct {
	Keyboard        [][]Keyboard `json:"keyboard"`
	OneTimeKeyboard bool         `json:"one_time_keyboard"`
	ResizeKeyboard  bool         `json:"resize_keyboard"`
}

type Keyboard struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}
