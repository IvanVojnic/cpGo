package service

import (
	"context"
	"fmt"
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
	"github.com/segmentio/kafka-go"
	"time"
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

func (s *UserCommService) SendInvite(userSender int, friendsList []int, id_place int) (string, error) {
	fmt.Println(id_place)
	message, err := s.repo.SendInvite(userSender, friendsList, id_place)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (s *UserCommService) GetRooms(userId int) ([]models.Rooms, error) {
	rooms, err := s.repo.GetRooms(userId)
	if err != nil {
		return rooms, err
	}
	errSend := sendRoomsPlaceId(rooms)
	if errSend != nil {
		return nil, errSend
	}
	conn, errGet := kafka.DialLeader(context.Background(), "tcp", "localhost", "quickstart-events", 0)
	if errGet != nil {
		return rooms, errGet
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 20))
	message, _ := conn.ReadMessage(1e6)
	for i := 0; i < len(rooms); i++ {
		rooms[i].PlaceName = string(message.Value)
	}
	return rooms, nil
}

func sendRoomsPlaceId(rooms []models.Rooms) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost", "placeConn", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	for i := 0; i < len(rooms); i++ {
		conn.WriteMessages(kafka.Message{Value: []byte(string(rooms[i].Place))})
	}
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

/*func getRoomsPlaceName(rooms *[]models.Rooms) error {

	if err != nil {
		return err
	}
	conn.Close()
	return nil
}
*/
