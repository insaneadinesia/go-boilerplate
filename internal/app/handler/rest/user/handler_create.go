package user

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/response"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

// @Tags Users
// @Summary Create User
// @Description API for create user
// @Accept json
// @Produce json
// @Param payload body user.CreateUpdateUserRequest true "Payload Create User"
// @Success 200 {object} response.DefaultResponse{data=nil}
// @Failure 400 {object} response.ErrorResponse{data=nil}
// @Failure 422 {object} response.ErrorResponse{data=nil}
// @Failure 500 {object} response.ErrorResponse{data=nil}
// @Router /users [post]
func (h *handler) Create(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := user.CreateUpdateUserRequest{}
	if err = validator.Validate(c, &req); err != nil {
		return
	}

	err = h.userUsecase.Create(ctx, req)
	if err != nil {
		return
	}

	return response.Success(c, nil)
}
