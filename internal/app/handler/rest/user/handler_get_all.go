package user

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/response"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

// @Tags Users
// @Summary Get All Users
// @Description API for get all users
// @Produce json
// @Param page query int false "page" default(1)
// @Param per_page query int false "per_page" default(20)
// @Param name query string false "search by name"
// @Param username query string false "search by username"
// @Param email query string false "search by email"
// @Success 200 {object} response.DefaultResponse{data=user.GetAllUserResponse}
// @Failure 400 {object} response.ErrorResponse{data=nil}
// @Failure 500 {object} response.ErrorResponse{data=nil}
// @Router /users [get]
func (h *handler) GetAll(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := user.GetAllUserRequest{}
	if err = validator.Validate(c, &req); err != nil {
		return
	}
	resp, err := h.userUsecase.GetAll(ctx, req)
	if err != nil {
		return
	}

	return response.Success(c, resp)
}
