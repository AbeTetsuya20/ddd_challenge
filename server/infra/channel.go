package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
)

type ChannelRepository struct {
	Conn *sql.Conn
}

func (c ChannelRepository) CreateChannel(ctx context.Context, channel *domain.Channel) error {
	//TODO implement me
	panic("implement me")
}

func (c ChannelRepository) GetChannel(ctx context.Context) (*domain.Channel, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChannelRepository) UpdateChannel(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error {
	//TODO implement me
	panic("implement me")
}

func (c ChannelRepository) DeleteChannel(ctx context.Context, channel *domain.Channel) error {
	//TODO implement me
	panic("implement me")
}
