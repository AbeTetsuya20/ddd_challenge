package model

import (
	"fmt"
	"time"
)

type ChannelID string

func NewChannelID(uuid int) ChannelID {
	return ChannelID(fmt.Sprintf("channelID_%d", uuid))
}

type Channel struct {
	ChannelID   ChannelID
	ChannelName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
