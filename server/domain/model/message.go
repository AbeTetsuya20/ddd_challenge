package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type MessageID string

func NewMessageID(messageTitle string) MessageID {
	uuid, _ := uuid.NewRandom()
	return MessageID(fmt.Sprintf("MessageID_%s_%s", messageTitle, uuid))
}

type Message struct {
	MessageID   MessageID
	MessageBody string
	Author      UserID
	ChannelID   ChannelID
	IsSend      bool
	SendAt      time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
