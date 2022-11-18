package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
)

type UserRepository struct {
	Conn *sql.Conn
}

func (u UserRepository) CreateUser(ctx context.Context, user *domain.User) error {

}

func (u UserRepository) GetUser(ctx context.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) UpdateUser(ctx context.Context, beforeUser *domain.User, afterUser *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) DeleteUser(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
