package model

import (
	"fmt"
	"time"
)

type MessageID string

func NewMessageID(messageTitle string) MessageID {
	return MessageID(fmt.Sprintf("messageID_%s", messageTitle))
}

type Message struct {
	MessageID    MessageID
	MessageTitle string
	MessageBody  string
	Author       UserID
	IsSend       bool
	SendAt       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewMessage(messageTitle, messageBody string, author UserID, sendAt time.Time) *Message {
	now := time.Now()
	return &Message{
		MessageID:    NewMessageID(messageTitle),
		MessageTitle: messageTitle,
		MessageBody:  messageBody,
		Author:       author,
		IsSend:       false,
		SendAt:       sendAt,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
