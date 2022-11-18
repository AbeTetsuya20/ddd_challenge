package handler

import (
	"context"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"github.com/AbeTetsuya20/ddd_challenge/server/usecase"
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
	message := &domain.Message{}
	s.Controller.SendMessage(ctx, message)
}
