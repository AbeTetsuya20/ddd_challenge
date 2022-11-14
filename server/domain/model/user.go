package model

import "time"

// User n <--> 1 Member
// User 1 <--> n Message
// Message n <--> 1 Member

type User struct {
	UserID    string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Member struct {
	MemberID    string
	Users       []*User
	MessageList []*Message
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Message struct {
	MessageID   string
	MessageBody string
	Author      User
	IsSend      bool
	SendAt      time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
