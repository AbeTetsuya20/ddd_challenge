package infra

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"net/http"
)

type UserRepository struct {
	conn *sql.Conn
	r    *http.Request
}

func NewUserRepository(conn *sql.Conn, r *http.Request) *UserRepository {
	return &UserRepository{conn: conn, r: r}
}

func (u UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	// Header から情報を取得
	headerName := u.r.Header.Get("name")
	headerPassword := u.r.Header.Get("password")

	// ログインチェック
	fmt.Println("headerName:", headerName)
	fmt.Println("headerPassword:", headerPassword)

	// 新規ユーザーの登録
	query := "INSERT INTO user () VALUES ()"
	_, err := u.conn.ExecContext(ctx, query)
	if err != nil {
		return errors.New(fmt.Sprintf("[ERROR] Insert: %+v", err))
	}

	return nil
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
