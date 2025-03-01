package user

import (
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/response"
	"github.com/labstack/echo/v4"
)

// @Tags Users
// @Summary Delete User By UUID
// @Description API for delete user by uuid
// @Produce json
// @Param uuid path string true "user uuid" default(d1e7cbc6-b6db-4f1f-a257-c6985dc2c2e3)
// @Success 200 {object} response.DefaultResponse{data=nil}
// @Failure 404 {object} response.ErrorResponse{data=nil}
// @Failure 422 {object} response.ErrorResponse{data=nil}
// @Failure 500 {object} response.ErrorResponse{data=nil}
// @Router /users/{uuid} [delete]
func (h *handler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	uuid := c.Param("uuid")
	err = h.userUsecase.Delete(ctx, uuid)
	if err != nil {
		return
	}

	return response.Success(c, nil)
}
