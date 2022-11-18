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
	Channel      ChannelID
	IsSend       bool
	SendAt       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
