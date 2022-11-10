package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
)

const salt = "s53d42fg98gh7j6kkbver"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
