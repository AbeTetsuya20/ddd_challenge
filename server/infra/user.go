package infra

import (
	"context"
	"database/sql"
	"fmt"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"log"
	"net/http"
)

type UserRepository struct {
	Conn *sql.Conn
	r    *http.Request
}

func NewUserRepository(conn *sql.Conn, r *http.Request) *UserRepository {
	return &UserRepository{Conn: conn, r: r}
}

func (u UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	// Header から情報を取得
	headerName := u.r.Header.Get("name")
	headerPassword := u.r.Header.Get("password")

	// ログインチェック
	fmt.Println("headerName:", headerName)
	fmt.Println("headerPassword:", headerPassword)

	// 新規ユーザーの登録
	query := "INSERT INTO user (user_ID, user_name, created_at, updated_at) VALUES (?,?,?,?) "
	_, err := u.Conn.ExecContext(ctx, query, user.UserID, user.UserName, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create User: %+v", err)
		return nil
	}

	return nil
}

func (u UserRepository) GetUser(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	query := "SELECT * FROM user WHERE UserID = ?"
	rows, err := u.Conn.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("[ERROR] can't create User: %+v", err)
		return nil, err
	}
	// todo ここから
}

func (u UserRepository) UpdateUser(ctx context.Context, userID string, updatedUser *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) DeleteUser(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
