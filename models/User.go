package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Name     string `json:"name" db:"name" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
