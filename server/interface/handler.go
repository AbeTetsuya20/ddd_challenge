package handler

import (
	"context"
	"encoding/json"
	"fmt"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"github.com/AbeTetsuya20/ddd_challenge/server/usecase"
	"log"
	"net/http"
)

type ServiceDriver struct {
	Controller usecase.Usecase
}

func NewServiceDriver(controller *usecase.ChatToolUsecase) *ServiceDriver {
	return &ServiceDriver{
		Controller: controller,
	}
}

func (s *ServiceDriver) MessageGetSend(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// channelID を header から取得
	channelID := r.Header.Get("channelID")

	// channelID が空文字だったらエラーを返す
	if channelID == "" {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", "channelID の情報が不足しています。"))
	}

	messages, err := s.Controller.MessageList(ctx, domain.ChannelID(channelID))
	if err != nil {
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%v", err))
	}

	// レスポンスを返す
	res := struct {
		messages []*domain.Message `json:"messages"`
	}{
		messages: messages,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		domain.WriteErrorResponse(w, 500, fmt.Sprintf("%w", err))
	}
}

func (s *ServiceDriver) MessageGetNotSend(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	message := &domain.Message{}
	s.Controller.SendMessage(ctx, message)
}
