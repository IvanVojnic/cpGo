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
	query := fmt.Sprintf("UPDATE FRIENDS SET FRIENDS.status = 'friends' WHERE FRIENDS.userSender = $1 AND userReceiver = $2")
	result, err := r.db.Exec(query, userSender, userReceiver)
	if err != nil {
		return "error", err
	}
	fmt.Println(result)
	return "accepted", nil
}
