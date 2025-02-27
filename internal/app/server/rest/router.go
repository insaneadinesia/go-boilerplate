package rest

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/app/handler/rest/health_check"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRouter(server *echo.Echo, container *container.Container) {
	// inject handler with usecase via container
	healthCheckHandler := health_check.NewHandler().SetHealthCheckUsecase(container.HealthCheckUsecase).Validate()

	server.GET("/health", healthCheckHandler.HealthCheck)
	server.GET("/swagger/*", echoSwagger.WrapHandler)
}
