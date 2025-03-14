package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
)

type handler struct {
	UnimplementedUserServiceServer
	userUsecase user.UserUsecase
	validator   *validator.Validate
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetUserUsecase(usecase user.UserUsecase) *handler {
	h.userUsecase = usecase
	return h
}

func (h *handler) SetValidator(v *validator.Validate) *handler {
	h.validator = v
	return h
}

func (h *handler) Validate() UserServiceServer {
	if h.userUsecase == nil {
		panic("userUsecase is nil")
	}

	if h.validator == nil {
		panic("validator is nil")
	}

	return h
}
