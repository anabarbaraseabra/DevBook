package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The field name is required")
	}
	if user.Nick == "" {
		return errors.New("The field nick is required")
	}
	if user.Email == "" {
		return errors.New("The field email is required")
	}
	if user.Password == "" {
		return errors.New("The field password is required")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Password = strings.TrimSpace(user.Password)
}
