package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
)

type JoinChannelToUserRepository struct {
	Conn *sql.Conn
}

func (j JoinChannelToUserRepository) GetChannelIDByUserID(ctx context.Context, userID domain.UserID) ([]domain.ChannelID, error) {
	//TODO implement me
	panic("implement me")
}

func (j JoinChannelToUserRepository) GetUserIDByChannelID(ctx context.Context, userID domain.ChannelID) ([]domain.UserID, error) {
	//TODO implement me
	panic("implement me")
}

func (j JoinChannelToUserRepository) CreateConnectionUserIDToChannelID(ctx context.Context, userid domain.UserID, channelID domain.ChannelID) error {
	//TODO implement me
	panic("implement me")
}

func (j JoinChannelToUserRepository) DeleteConnectionUserIDToChannelID(ctx context.Context, userid domain.UserID, channelID domain.ChannelID) error {
	//TODO implement me
	panic("implement me")
}
