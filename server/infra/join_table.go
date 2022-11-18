package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"log"
)

type JoinChannelToUserRepository struct {
	Conn *sql.Conn
}

func ScanChannelIDs(rows *sql.Rows) ([]domain.ChannelID, int, error) {
	channelIDs := make([]domain.ChannelID, 0)

	for rows.Next() {
		var v domain.ChannelID
		if err := rows.Scan(&v); err != nil {
			log.Printf("[ERROR] scan ScanChannelIDs: %+v", err)
			return nil, 0, err
		}
		channelIDs = append(channelIDs, v)
	}

	return channelIDs, len(channelIDs), nil
}

func ScanUserIDs(rows *sql.Rows) ([]domain.UserID, int, error) {
	userIDs := make([]domain.UserID, 0)

	for rows.Next() {
		var v domain.UserID
		if err := rows.Scan(&v); err != nil {
			log.Printf("[ERROR] scan ScanUserIDs: %+v", err)
			return nil, 0, err
		}
		userIDs = append(userIDs, v)
	}

	return userIDs, len(userIDs), nil
}

func (j JoinChannelToUserRepository) GetChannelIDsByUserID(ctx context.Context, userID domain.UserID) ([]domain.ChannelID, error) {
	query := "SELECT * FROM joinChannelToUser WHERE UserID = ?"
	rows, err := j.Conn.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("[ERROR] can't get GetChannelIDsByUserID: %+v", err)
		return nil, err
	}

	channels, _, err := ScanChannelIDs(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan Channels: %+v", err)
		return nil, err
	}

	return channels, nil
}

func (j JoinChannelToUserRepository) GetUserIDsByChannelID(ctx context.Context, channelID domain.ChannelID) ([]domain.UserID, error) {
	query := "SELECT * FROM joinChannelToUser WHERE ChannelID = ?"
	rows, err := j.Conn.QueryContext(ctx, query, channelID)
	if err != nil {
		log.Printf("[ERROR] can't get GetUserIDsByChannelID: %+v", err)
		return nil, err
	}

	userIDs, _, err := ScanUserIDs(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan Channels: %+v", err)
		return nil, err
	}

	return userIDs, nil
}

func (j JoinChannelToUserRepository) CreateConnectionUserIDToChannelID(ctx context.Context, join *domain.JoinChannelToUser) error {
	query := "INSERT INTO joinChannelToUser (user_ID, user_name, channel_ID, channel_name, created_at, updated_at) VALUES (?,?,?,?,?,?) "
	_, err := j.Conn.ExecContext(ctx, query, join.UserID, join.UserName, join.ChannelID, join.ChannelName, join.CreatedAt, join.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create CreateConnectionUserIDToChannelID: %+v", err)
		return nil
	}

	return nil
}

func (j JoinChannelToUserRepository) DeleteConnectionUserIDToChannelID(ctx context.Context, userid domain.UserID, channelID domain.ChannelID) error {
	query := "DELETE FROM channel WHERE user_id = ? AND channel_id = ?"
	_, err := j.Conn.ExecContext(ctx, query, userid, channelID)
	if err != nil {
		log.Printf("[ERROR] can't delete DeleteConnectionUserIDToChannelID: %+v", err)
		return nil
	}

	return nil
}
