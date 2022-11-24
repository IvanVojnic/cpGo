package repository

import (
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type UserCommunicate interface {
	FindUserByEmail(email string) (models.User, error)
	SendRequestToFriends(userSender int, userReceiver int) error
	GetFriendsRequestById(userId int) ([]models.User, error)
	AcceptFriendsRequest(userSender, userReceiver int) (string, error)
}

type Repository struct {
	Authorization
	UserCommunicate
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		UserCommunicate: NewUserCommPostgres(db),
	}
}
