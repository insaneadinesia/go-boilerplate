package user

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Create(c echo.Context) (err error)
	GetAll(c echo.Context) (err error)
	GetDetail(c echo.Context) (err error)
	Update(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
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

func (h *handler) Validate() UserHandler {
	if h.userUsecase == nil {
		panic("userUsecase is nil")
	}

	return h
}
