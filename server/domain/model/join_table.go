package model

import "time"

type JoinChannelToUser struct {
	UserID      UserID
	UserName    string
	ChannelID   ChannelID
	ChannelName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
