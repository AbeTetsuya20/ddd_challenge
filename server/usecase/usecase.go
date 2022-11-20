package usecase

import (
	"context"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"github.com/AbeTetsuya20/ddd_challenge/server/domain/repository"
)

type Usecase interface {

	// メッセージに関する UseCase

	// ユーザーは、【チャット】の【所属メンバー】に【メッセージ】を送信する
	CreateMessage(ctx context.Context, message *domain.Message) error
	GetMessageByIsSend(ctx context.Context, channelID domain.ChannelID) ([]domain.Message, error)
	// ユーザーは、【送信予定のメッセージ】を全て確認できる
	GetMessageByNotIsSend(ctx context.Context, channelID domain.ChannelID, userID domain.UserID) ([]domain.Message, error)
	// ユーザーは、【送信予定のメッセージ】を編集・削除できる
	UpdateMessageByNotIsSend(ctx context.Context, afterMessage *domain.Message) error

	// ユーザーに関する UseCase

	// User を作成する
	CreateUser(ctx context.Context, user *domain.User) error
	// ChannelID でユーザーを取得する

	// チャンネルに関する UseCase

	// チャンネルを作成する
	CreateChannel(ctx context.Context, channel *domain.Channel) error
	// ユーザーは、【チャット】の【名前】,【所属メンバー】などを設定する。
	EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error
	// 全チャンネルを取得する
	GetChannels(ctx context.Context) ([]domain.Channel, error)
	// UserID でチャンネルを取得する
	GetChannelByUserID(ctx context.Context, userID domain.UserID) ([]*domain.JoinChannelToUser, error)
	// チャンネルを削除する
	DeleteChannel(ctx context.Context, channelD domain.ChannelID) error
}

type ChatToolUsecase struct {
	UserRepo    repository.UserRepository
	ChannelRepo repository.ChannelRepository
	MessageRepo repository.MessageRepository
	JoinRepo    repository.JoinChannelToUserRepository
}

func NewChatToolUsecase(userRepo repository.UserRepository, channelRepo repository.ChannelRepository, messageRepo repository.MessageRepository, joinRepo repository.JoinChannelToUserRepository) *ChatToolUsecase {
	return &ChatToolUsecase{UserRepo: userRepo, ChannelRepo: channelRepo, MessageRepo: messageRepo, JoinRepo: joinRepo}
}

func (c *ChatToolUsecase) CreateMessage(ctx context.Context, message *domain.Message) error {
	err := c.MessageRepo.CreateMessage(ctx, message)
	return err
}

func (c *ChatToolUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	err := c.UserRepo.CreateUser(ctx, user)
	return err
}

func (c *ChatToolUsecase) CreateChannel(ctx context.Context, channel *domain.Channel) error {
	err := c.ChannelRepo.CreateChannel(ctx, channel)
	return err
}

func (c *ChatToolUsecase) EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error {
	err := c.ChannelRepo.UpdateChannel(ctx, "", afterChannel)
	return err
}

func (c *ChatToolUsecase) DeleteChannel(ctx context.Context, channelD domain.ChannelID) error {
	err := c.ChannelRepo.DeleteChannel(ctx, channelD)
	return err
}

func (c *ChatToolUsecase) GetMessageByIsSend(ctx context.Context, channelID domain.ChannelID) ([]domain.Message, error) {
	messages, err := c.MessageRepo.GetAllSendMessages(ctx, channelID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (c *ChatToolUsecase) GetMessageByNotIsSend(ctx context.Context, channelID domain.ChannelID, userID domain.UserID) ([]domain.Message, error) {
	messages, err := c.MessageRepo.GetMessagesByChannelIDAndIsNotSendAndUserID(ctx, channelID, userID)
	if err != nil {
		return nil, err
	}
	return messages, nil

}

func (c *ChatToolUsecase) GetChannels(ctx context.Context) ([]domain.Channel, error) {
	channel, err := c.ChannelRepo.GetChannels(ctx)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (c *ChatToolUsecase) GetChannelByUserID(ctx context.Context, userID domain.UserID) ([]*domain.JoinChannelToUser, error) {
	joinChannelToUsers, err := c.JoinRepo.GetChannelIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return joinChannelToUsers, nil
}

func (c *ChatToolUsecase) UpdateMessageByNotIsSend(ctx context.Context, afterMessage *domain.Message) error {
	err := c.MessageRepo.UpdateMessage(ctx, afterMessage)
	return err
}
