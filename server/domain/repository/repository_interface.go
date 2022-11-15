package repository

import (
	"context"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context) (*domain.User, error)
	UpdateUser(ctx context.Context, beforeUser *domain.User, afterUser *domain.User) error
	DeleteUser(ctx context.Context, user *domain.User) error
}

type ChannelRepository interface {
	CreateChannel(ctx context.Context, channel *domain.Channel) error
	GetChannel(ctx context.Context) (*domain.Channel, error)
	UpdateChannel(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error
	DeleteChannel(ctx context.Context, channel *domain.Channel) error
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *domain.Message) error
	GetMessage(ctx context.Context) (*domain.Message, error)
	UpdateMessage(ctx context.Context, beforeMessage *domain.Message, afterMessage *domain.Message) error
	DeleteMessage(ctx context.Context, message *domain.Message) error
}
