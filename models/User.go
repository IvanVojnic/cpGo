package models

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"'`
}
