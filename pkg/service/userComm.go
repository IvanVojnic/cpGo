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
	/*user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))*/
}
