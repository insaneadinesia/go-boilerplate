package user

import (
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/response"
	"github.com/labstack/echo/v4"
)

// @Tags Users
// @Summary Get User Detail By UUID
// @Description API for get user detail by uuid
// @Produce json
// @Param uuid path string true "user uuid" default(d1e7cbc6-b6db-4f1f-a257-c6985dc2c2e3)
// @Success 200 {object} response.DefaultResponse{data=user.UserDetailResponse}
// @Failure 404 {object} response.ErrorResponse{data=nil}
// @Failure 400 {object} response.ErrorResponse{data=nil}
// @Failure 500 {object} response.ErrorResponse{data=nil}
// @Router /users/{uuid} [get]
func (h *handler) GetDetail(c echo.Context) (err error) {
	ctx := c.Request().Context()

	uuid := c.Param("uuid")
	resp, err := h.userUsecase.GetDetail(ctx, uuid)
	if err != nil {
		return
	}

	return response.Success(c, resp)
}
