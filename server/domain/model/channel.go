package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ChannelID string

func NewChannelID() ChannelID {
	uuid, _ := uuid.NewRandom()
	return ChannelID(fmt.Sprintf("ChannelID_%s", uuid))
}

type Channel struct {
	ChannelID   ChannelID
	ChannelName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
