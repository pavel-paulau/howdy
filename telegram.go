package main

// User object represents a Telegram user or bot.
//
// https://core.telegram.org/bots/api#user
//
type User struct {
	ID        int    `json:"id"`         // Unique identifier for this user or bot
	FirstName string `json:"first_name"` // User‘s or bot’s first name
}

// Chat object represents a chat.
//
// https://core.telegram.org/bots/api#chat
//
type Chat struct {
	ID   int    `json:"id"`   // Unique identifier for this chat.
	Type string `json:"type"` // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
}

// Contact object represents a phone contact.
//
// https://core.telegram.org/bots/api#contact
//
type Contact struct {
	FirstName   string `json:"first_name"`   // Contact's phone number
	PhoneNumber string `json:"phone_number"` // Contact's first name
}

// Message object represents a message.
//
// https://core.telegram.org/bots/api#message
//
type Message struct {
	MessageID int     `json:"message_id"` // Unique message identifier
	From      User    `json:"from"`       // Sender
	Chat      Chat    `json:"chat"`       // Conversation the message belongs to
	Date      int     `json:"date"`       // Date the message was sent in Unix time
	Text      string  `json:"text"`       // The actual UTF-8 text of the message
	Contact   Contact `json:"contact"`    // Information about the contact
}

// Update object represents an incoming update.
//
// https://core.telegram.org/bots/api#update
//
type Update struct {
	UpdateID int     `json:"update_id"` // The update‘s unique identifier
	Message  Message `json:"message"`   // New incoming message of any kind
}

// ReplyKeyboardMarkup object represents a custom keyboard with reply options
//
// https://core.telegram.org/bots/api#replykeyboardmarkup
//
type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"` // Array of button rows
}

// KeyboardButton object represents one button of the reply keyboard.
//
// https://core.telegram.org/bots/api#keyboardbutton
//
type KeyboardButton struct {
	Text           string `json:"text"`            // Text of the button
	RequestContact bool   `json:"request_contact"` // If True, the user's phone number will be sent as a contact
}

type telegramResponse struct {
	OK bool `json:"ok"`
}

type botResponse struct {
	ChatID      int         `json:"chat_id"`      // Unique identifier for the target chat
	Text        string      `json:"text"`         // Text of the message to be sent
	ReplyMarkup interface{} `json:"reply_markup"` // Additional interface options.
}
