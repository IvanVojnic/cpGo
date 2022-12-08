package models

type Rooms struct {
	Id            int    `json:"id" db:"id"`
	IdUserCreator int    `json:"id_user_creator" db:"id_user_creator" binding:"required"`
	DateEvent     string `json:"date_event" db:"date_event"`
	Place         int    `json:"id_place" db:"id_place" binding:"required"`
	PlaceName     string `json:"place_name"`
}
