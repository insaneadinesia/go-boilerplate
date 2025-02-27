package health_check

import (
	"net/http"

	"github.com/insaneadinesia/go-boilerplate/internal/pkg/response"
	"github.com/labstack/echo/v4"
)

// @Summary Health Check
// @Description API for do health check to dependencies used
// @Produce json
// @Success 200 {object} response.DefaultResponse{data=health_check.StatusCheck}
// @Failure 503 {object} response.DefaultResponse{data=health_check.StatusCheck}
// @Router /health [get]
func (h *handler) HealthCheck(c echo.Context) error {
	ctx := c.Request().Context()

	code := http.StatusOK
	resp, err := h.healthCheckUsecase.HealthCheck(ctx)
	if err != nil {
		code = http.StatusServiceUnavailable
	}

	return response.Default(c, code, resp)
}
