package service

import (
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUser(email, password string) (int, error)
}

type UserCommunicate interface {
	FindUser(email string) (models.User, error)
	SendRequest(userSender int, userReceiver int) error
	GetFriendsRequest(userId int) ([]models.User, error)
	AcceptFriendsRequest(userSender, userReceiver int) (string, error)
}

type Service struct {
	Authorization
	UserCommunicate
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repos.Authorization),
		UserCommunicate: NewUserComm(repos.UserCommunicate),
	}
}
