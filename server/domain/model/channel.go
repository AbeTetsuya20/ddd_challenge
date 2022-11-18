package model

import (
	"fmt"
	"math/rand"
	"time"
)

type ChannelID string

func NewChannelID(uuid int) ChannelID {
	return ChannelID(fmt.Sprintf("channelID_%d", uuid))
}

type Channel struct {
	ChannelID   ChannelID
	Users       []*UserID
	MessageList []*MessageID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewChannel(users []*UserID, messages []*MessageID) *Channel {
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	uuid := rand.Intn(100)
	return &Channel{
		ChannelID:   NewChannelID(uuid),
		Users:       users,
		MessageList: messages,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
