package repository

import (
	"fmt"
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/jmoiron/sqlx"
)

type UserCommPostgres struct {
	db *sqlx.DB
}

func NewUserCommPostgres(db *sqlx.DB) *UserCommPostgres {
	return &UserCommPostgres{db: db}
}

func (r *UserCommPostgres) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, name, email FROM %s WHERE email=$1", usersTable)
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *UserCommPostgres) SendRequestToFriends(userSender int, userReceiver int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_sender, user_receiver, status) values ($1, $2, $3)", friendsTable)
	row := r.db.QueryRow(query, userSender, userReceiver, "request")
	if err := row.Scan(&userSender); err != nil {
		return err
	}
	return nil
}

func (r *UserCommPostgres) GetFriendsRequestById(userId int) ([]models.User, error) {
	var users1 []models.User
	query := fmt.Sprintf("SELECT USERS.id, USERS.email, USERS.name FROM USERS INNER JOIN FRIENDS on FRIENDS.user_receiver = USERS.id WHERE FRIENDS.user_sender = $1 AND FRIENDS.status = 'request'")
	err := r.db.Select(&users1, query, userId)
	if err != nil {
		fmt.Println(err)
	}
	var users2 []models.User
	query2 := fmt.Sprintf("SELECT USERS.id, USERS.email, USERS.name FROM USERS INNER JOIN FRIENDS on FRIENDS.user_sender = USERS.id WHERE FRIENDS.user_receiver = $1 AND FRIENDS.status = 'request'")
	err2 := r.db.Select(&users2, query2, userId)
	var allUsers []models.User
	allUsers = append(users1, users2...)
	return allUsers, err2
}

func (r *UserCommPostgres) AcceptFriendsRequest(userSender, userReceiver int) (string, error) {
	query := fmt.Sprintf("UPDATE FRIENDS SET status = 'friends' WHERE user_sender = $1 AND user_receiver = $2")
	_, err := r.db.Exec(query, userSender, userReceiver)
	if err != nil {
		return "error", err
	}
	return "accepted", nil
}

func (r *UserCommPostgres) GetAllFriends(userId int) ([]models.User, error) {
	var users1 []models.User
	query := fmt.Sprintf("SELECT USERS.id, USERS.email, USERS.name FROM USERS INNER JOIN FRIENDS on FRIENDS.user_receiver = USERS.id WHERE FRIENDS.user_sender = $1 AND FRIENDS.status = 'friends'")
	err := r.db.Select(&users1, query, userId)
	if err != nil {
		fmt.Println(err)
	}
	var users2 []models.User
	query2 := fmt.Sprintf("SELECT USERS.id, USERS.email, USERS.name FROM USERS INNER JOIN FRIENDS on FRIENDS.user_sender = USERS.id WHERE FRIENDS.user_receiver = $1 AND FRIENDS.status = 'friends'")
	err2 := r.db.Select(&users2, query2, userId)
	var allUsers []models.User
	allUsers = append(users1, users2...)
	return allUsers, err2
}

func (r *UserCommPostgres) SendInvite(userSender int, friendsList []int) (string, error) {
	var roomId int
	queryRoom := fmt.Sprintf("INSERT INTO %s (id_user_creator, date_event, place) values ($1, $2, $3) RETURNING id", roomsTable)
	rowRoom := r.db.QueryRow(queryRoom, userSender, "2022-10-10", "place")
	if err := rowRoom.Scan(&roomId); err != nil {
		return "error", err
	}
	fmt.Println("room id")
	fmt.Println(roomId)
	for i := 0; i < len(friendsList); i++ {
		query := fmt.Sprintf("INSERT INTO %s (status_id, user_id, room_id) values ($1, $2, $3) RETURNING id", invitesTable)
		row := r.db.QueryRow(query, 2, friendsList[i], roomId)
		if err := row.Scan(&roomId); err != nil {
			return "error invites", err
		}
	}
	return "invite sends", nil
}

func (r *UserCommPostgres) GetRooms(userId int) ([]models.Rooms, error) {
	fmt.Println("getRooms 2")
	fmt.Println(userId)
	var rooms1 []models.Rooms
	//query1 := fmt.Sprintf("SELECT ROOMS.id, ROOMS.id_user_creator, ROOMS.place FROM ROOMS WHERE ROOMS.id_user_creator = $1")
	query1 := fmt.Sprintf("SELECT id, id_user_creator, date_event, place FROM %s WHERE id_user_creator=$1", roomsTable)
	err1 := r.db.Select(&rooms1, query1, userId)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("getRooms 3")
	fmt.Println(rooms1)
	var rooms2 []models.Rooms
	query2 := fmt.Sprintf("SELECT ROOMS.id, ROOMS.id_user_creator, ROOMS.date_event, ROOMS.place FROM ROOMS INNER JOIN INVITES on INVITES.room_id = ROOMS.id WHERE INVITES.user_id = $1")
	err2 := r.db.Select(&rooms2, query2, userId)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("getRooms 4")
	fmt.Println(rooms2)
	var allRooms []models.Rooms
	allRooms = append(rooms1, rooms2...)
	return allRooms, nil
}
