package models

type Friends struct {
	UserSender   int    `json:"user_sender" db:"user_sender"`
	UserReceiver int    `json:"user_receiver" db:"user_receiver"`
	Status       string `json:"status" db:"status"`
}
