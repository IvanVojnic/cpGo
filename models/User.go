package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
