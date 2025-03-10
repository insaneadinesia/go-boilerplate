package user

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
)

type WorkerUserHandler interface {
	InformUserCreated(ctx context.Context, t *asynq.Task) (err error)
}

type handler struct {
	userUsecase user.UserUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetUserUsecase(usecase user.UserUsecase) *handler {
	h.userUsecase = usecase
	return h
}

func (h *handler) Validate() WorkerUserHandler {
	if h.userUsecase == nil {
		panic("userUsecase is nil")
	}

	return h
}

func (h *handler) InformUserCreated(ctx context.Context, t *asynq.Task) (err error) {
	req := user.CreatedUserPayload{}
	if err = json.Unmarshal(t.Payload(), &req); err != nil {
		return
	}

	if err = h.userUsecase.Inform3rdParty(ctx, req); err != nil {
		return
	}

	return
}
