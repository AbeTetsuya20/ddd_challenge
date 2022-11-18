package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
)

type MessageRepository struct {
	Conn *sql.Conn
}

func NewMessageRepository(conn *sql.Conn) *MessageRepository {
	return &MessageRepository{Conn: conn}
}

func (m MessageRepository) CreateMessage(ctx context.Context, message *domain.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) GetAllSendMessage(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) GetMessageByChannelIDByIsNotSendAndUserID(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) UpdateMessage(ctx context.Context, afterMessage *domain.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) DeleteMessage(ctx context.Context, message *domain.Message) error {
	//TODO implement me
	panic("implement me")
}
