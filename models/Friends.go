package models

type Friends struct {
	UserSender   int    `json:"-" db:"userSender"`
	UserReceiver int    `json:"-" db:"userReceiver"`
	status       string `json:"-" db:"status"`
}
