package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"log"
)

type UserRepository struct {
	Conn *sql.Conn
}

func NewUserRepository(conn *sql.Conn) *UserRepository {
	return &UserRepository{Conn: conn}
}

func (u UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	// 新規ユーザーの登録
	query := "INSERT INTO user (user_ID, user_name, user_password, created_at, updated_at) VALUES (?,?,?,?,?) "
	_, err := u.Conn.ExecContext(ctx, query, user.UserID, user.UserName, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create User: %+v", err)
		return nil
	}

	return nil
}

func (u UserRepository) GetUser(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	query := "SELECT * FROM user WHERE user_id = ?"
	rows, err := u.Conn.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("[ERROR] can't get User: %+v", err)
		return nil, err
	}

	var user *domain.User
	if err := rows.Scan(user); err != nil {
		log.Printf("[ERROR] scan ScanChannels: %+v", err)
		return nil, err
	}

	return user, nil
}

func (u UserRepository) UpdateUser(ctx context.Context, userID string, updatedUser *domain.User) error {
	query := "UPDATE user set user_name = ? WHERE user_id = ? "
	_, err := u.Conn.ExecContext(ctx, query, updatedUser, userID)
	if err != nil {
		log.Printf("[ERROR] can't UpdateUser: %+v", err)
		return nil
	}

	return nil
}

func (u UserRepository) DeleteUser(ctx context.Context, userID domain.UserID) error {
	query := "DELETE FROM user WHERE user_id = ?"
	_, err := u.Conn.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("[ERROR] can't delete user: %+v", err)
		return nil
	}

	return nil
}
