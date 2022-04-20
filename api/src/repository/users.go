package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (u *users) SearchByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%
	rows, err := u.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?", //comparacao com like busca qualquer coisa que tenha o trecho da variavel
		nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u *users) GetByID(userID string) (models.User, error) {
	rows, err := u.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", userID)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()
	var user models.User
	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
