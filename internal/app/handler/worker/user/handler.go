package user

import (
	"context"
	"encoding/json"

	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
)

type WorkerUserHandler interface {
	InformUserCreated(ctx context.Context, payload any) (err error)
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

func (h *handler) InformUserCreated(ctx context.Context, payload any) (err error) {
	req := user.CreatedUserPayload{}

	by, _ := json.Marshal(payload)
	err = json.Unmarshal(by, &req)
	if err != nil {
		return
	}

	err = h.userUsecase.Inform3rdParty(ctx, req)
	if err != nil {
		return
	}

	return
}
