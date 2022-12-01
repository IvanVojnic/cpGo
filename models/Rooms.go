package models

type Rooms struct {
	Id            int    `json:"id" db:"id"`
	IdUserCreator int    `json:"id_user_creator" db:"id_user_creator" binding:"required"`
	DateEvent     string `json:"date_event" db:"date_event"`
	Place         string `json:"place" db:"place" binding:"required"`
}
