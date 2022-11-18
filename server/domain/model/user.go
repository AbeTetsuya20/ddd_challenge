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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(userName string) *User {
	now := time.Now()
	return &User{
		UserID:    NewUserID(userName),
		Name:      userName,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
