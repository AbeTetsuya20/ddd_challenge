package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"log"
)

type MessageRepository struct {
	Conn *sql.DB
}

func NewMessageRepository(conn *sql.DB) *MessageRepository {
	return &MessageRepository{Conn: conn}
}

func ScanMessages(rows *sql.Rows) ([]*domain.Message, int, error) {
	messages := make([]*domain.Message, 0)

	for rows.Next() {
		var v *domain.Message
		if err := rows.Scan(&v.ChannelID, &v.CreatedAt, &v.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan ScanChannels: %+v", err)
			return nil, 0, err
		}
		messages = append(messages, v)
	}

	return messages, len(messages), nil
}

func (m MessageRepository) CreateMessage(ctx context.Context, message *domain.Message) error {
	query := "INSERT INTO user (message_id, message_body, author, channel_id, is_send, send_at, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?) "
	_, err := m.Conn.ExecContext(ctx, query, message.MessageID, message.MessageBody, message.Author, message.ChannelID, message.IsSend, message.SendAt, message.CreatedAt, message.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create Message: %+v", err)
		return nil
	}

	return nil
}

func (m MessageRepository) GetAllSendMessage(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error) {
	query := "SELECT * FROM message WHERE ChannelID = ? AND is_send = true"
	rows, err := m.Conn.QueryContext(ctx, query, channelID)
	if err != nil {
		log.Printf("[ERROR] can't get GetAllSendMessage: %+v", err)
		return nil, err
	}

	messages, _, err := ScanMessages(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan Channels: %+v", err)
		return nil, err
	}

	return messages, nil
}

func (m MessageRepository) GetMessageByChannelIDByIsNotSendAndUserID(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) UpdateMessage(ctx context.Context, afterMessage *domain.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) DeleteMessage(ctx context.Context, message *domain.Message) error {
	//TODO implement me
	panic("implement me")
}
