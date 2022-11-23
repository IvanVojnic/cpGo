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
	var friends models.Friends
	query := fmt.Sprintf("INSERT INTO %s (userSender, userReceiver, status) values ($1, $2, $3)", friendsTable)
	row := r.db.QueryRow(query, friends.UserSender, friends.UserReceiver, 1)
	if err := row.Scan(&userSender); err != nil {
		return err
	}
	return nil
}
