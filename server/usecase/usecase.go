package usecase

import (
	"context"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"github.com/AbeTetsuya20/ddd_challenge/server/domain/repository"
)

type Usecase interface {

	// ユーザーは、【チャット】の【所属メンバー】に【メッセージ】を送信する
	SendMessage(ctx context.Context, messageID domain.MessageID) error

	// User に関する Usecase

	CreateUser(ctx context.Context, user *domain.User) error

	// Channel に関する Usecase

	// チャンネルを作成する
	CreateChannel(ctx context.Context, channel *domain.Channel) error

	// ユーザーは、【チャット】の【名前】,【所属メンバー】などを設定する。つまり、チャンネルの名前を編集する？
	EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error

	// チャンネルを削除する
	DeleteChannel(ctx context.Context, channel *domain.Channel) error

	// チャンネルの全メッセージを取得する。これが画面に表示される
	MessageList(ctx context.Context, channelID *domain.ChannelID) (*domain.Message, error)

	// Message に関する Usecase

	// ユーザーは、【送信予定のメッセージ】を全て確認できる
	GetScheduleToSendMessage(ctx context.Context, channelID *domain.ChannelID) (*domain.Channel, error)

	// ユーザーは、【送信予定のメッセージ】を編集・削除できる
	EditScheduleToSendMessage(ctx context.Context, messageID *domain.MessageID) error

	// ユーザーは、【送信予定のメッセージ】の【送信予定時刻】を変更できる
	EditScheduleToSendMessageTime(ctx context.Context, messageID *domain.MessageID) error
}

type ChatToolUsecase struct {
	UserRepo    repository.UserRepository
	ChannelRepo repository.ChannelRepository
	MessageRepo repository.MessageRepository
}

func (c ChatToolUsecase) SendMessage(ctx context.Context, messageID domain.MessageID) error {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) CreateChannel(ctx context.Context, channel *domain.Channel) error {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) DeleteChannel(ctx context.Context, channel *domain.Channel) error {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) MessageList(ctx context.Context, channelID *domain.ChannelID) (*domain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) GetScheduleToSendMessage(ctx context.Context, channelID *domain.ChannelID) (*domain.Channel, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) EditScheduleToSendMessage(ctx context.Context, messageID *domain.MessageID) error {
	//TODO implement me
	panic("implement me")
}

func (c ChatToolUsecase) EditScheduleToSendMessageTime(ctx context.Context, messageID *domain.MessageID) error {
	//TODO implement me
	panic("implement me")
}
