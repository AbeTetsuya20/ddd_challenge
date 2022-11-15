package model

import "time"

// User n <--> 1 Member
// User 1 <--> n Message
// Message n <--> 1 Member

type UserID string

type User struct {
	UserID    UserID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ChannelID string

type Channel struct {
	ChannelID   ChannelID
	Users       []*User
	MessageList []*Message
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MessageID string

type Message struct {
	MessageID    MessageID
	MessageTitle string
	MessageBody  string
	Author       User
	IsSend       bool
	SendAt       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
