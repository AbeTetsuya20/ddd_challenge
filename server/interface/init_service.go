package handler

import (
	"database/sql"
	"github.com/AbeTetsuya20/ddd_challenge/server/infra"
	"github.com/AbeTetsuya20/ddd_challenge/server/usecase"
)

func InitService(conn *sql.Conn) (Service, error) {
	userRepo := infra.NewUserRepository(conn)
	channelRepo := infra.NewChannelRepository(conn)
	messageRepo := infra.NewMessageRepository(conn)
	joinRepo := infra.NewJoinChannelToUserRepository(conn)

	usecaseInterface := usecase.NewChatToolUsecase(userRepo, channelRepo, messageRepo, joinRepo)
	service := NewServiceDriver(usecaseInterface)
	return service, nil
}
