package repository

import (
	"api/src/models"
	"database/sql"
	"github.com/google/uuid"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db: db}
}

func (u *users) Create(user models.User) (string, error) {
	user.ID = uuid.NewString()
	statement, err := u.db.Prepare("INSERT INTO users (id, name, nick, email, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer statement.Close()

	_, err = statement.Exec(user.ID, user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
