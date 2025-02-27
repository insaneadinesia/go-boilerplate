package health_check

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/health_check"
	"github.com/labstack/echo/v4"
)

type HealthCheckHandler interface {
	HealthCheck(c echo.Context) error
}

type handler struct {
	healthCheckUsecase health_check.HealthCheckUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetHealthCheckUsecase(usecase health_check.HealthCheckUsecase) *handler {
	h.healthCheckUsecase = usecase

	return h
}

func (h *handler) Validate() HealthCheckHandler {
	if h.healthCheckUsecase == nil {
		panic("healthCheckUsecase is nil")
	}

	return h
}
