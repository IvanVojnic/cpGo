package models

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"'`
}
