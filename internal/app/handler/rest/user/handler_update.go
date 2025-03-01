package user

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/response"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

// @Tags Users
// @Summary Update User
// @Description API for update user
// @Accept json
// @Produce json
// @Param uuid path string true "user uuid" default(d1e7cbc6-b6db-4f1f-a257-c6985dc2c2e3)
// @Param payload body user.CreateUpdateUserRequest true "Payload Update User"
// @Success 200 {object} response.DefaultResponse{data=nil}
// @Failure 400 {object} response.ErrorResponse{data=nil}
// @Failure 422 {object} response.ErrorResponse{data=nil}
// @Failure 500 {object} response.ErrorResponse{data=nil}
// @Router /users/{uuid} [put]
func (h *handler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	uuid := c.Param("uuid")
	req := user.CreateUpdateUserRequest{}
	if err = validator.Validate(c, &req); err != nil {
		return
	}

	err = h.userUsecase.Update(ctx, uuid, req)
	if err != nil {
		return
	}

	return response.Success(c, nil)
}
