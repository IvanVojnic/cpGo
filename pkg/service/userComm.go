package service

import (
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
)

type UserCommService struct {
	repo repository.UserCommunicate
}

func NewUserComm(repo repository.UserCommunicate) *UserCommService {
	return &UserCommService{repo: repo}
}

func (s *UserCommService) FindUser(email string) (models.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return user, nil
	}
	return user, nil
}

func (s *UserCommService) SendRequest(userSender int, userReceiver int) error {
	err := s.repo.SendRequestToFriends(userSender, userReceiver)
	if err != nil {
		return nil
	}
	return err
}
