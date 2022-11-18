package model

import (
	"fmt"
	"time"
)

type UserID string

func NewUserID(name string) UserID {
	return UserID(fmt.Sprintf("userID_%s", name))
}

type User struct {
	UserID    UserID
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(userName string, password string) *User {
	now := time.Now()
	return &User{
		UserID:    NewUserID(userName),
		Name:      userName,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
