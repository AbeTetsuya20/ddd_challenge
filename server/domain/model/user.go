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
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
