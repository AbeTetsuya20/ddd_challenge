package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type UserID string

func NewUserID(name string) UserID {
	uuid, _ := uuid.NewRandom()
	return UserID(fmt.Sprintf("userID_%s_%s", name, uuid))
}

type User struct {
	UserID    UserID
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
