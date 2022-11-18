package handler

import (
	"database/sql"
	"github.com/AbeTetsuya20/ddd_challenge/server/infra"
)

func InitService(conn *sql.Conn) (Service, error) {
	_ = infra.NewUserRepository(conn)
	return nil, nil
}
