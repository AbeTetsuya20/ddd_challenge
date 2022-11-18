package handler

import (
	"database/sql"
	"github.com/AbeTetsuya20/ddd_challenge/server/infra"
)

func InitService(conn *sql.Conn) (Service, error) {
	userRepo := infra.NewUserRepository(conn)
	channelRepo := infra.NewChannelRepository(conn)
	messageRepo := infra.NewMessageRepository(conn)

	return nil, nil
}
