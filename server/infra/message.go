package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
)

type MessageRepository struct {
	Conn *sql.Conn
}

func (m MessageRepository) CreateMessage(ctx context.Context, message *domain.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) GetMessage(ctx context.Context) (*domain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) UpdateMessage(ctx context.Context, beforeMessage *domain.Message, afterMessage *domain.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) DeleteMessage(ctx context.Context, message *domain.Message) error {
	//TODO implement me
	panic("implement me")
}
