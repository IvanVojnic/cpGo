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
		return user, err
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

func (s *UserCommService) GetFriendsRequest(userId int) ([]models.User, error) {
	users, err := s.repo.GetFriendsRequestById(userId)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *UserCommService) AcceptFriendsRequest(userSender, userReceiver int) (string, error) {
	message, err := s.repo.AcceptFriendsRequest(userSender, userReceiver)
	if err != nil {
		return "error in accept", err
	}
	return message, nil
}

func (s *UserCommService) GetAllFriends(userId int) ([]models.User, error) {
	users, err := s.repo.GetAllFriends(userId)
	if err != nil {
		return users, err
	}
	return users, nil
}
