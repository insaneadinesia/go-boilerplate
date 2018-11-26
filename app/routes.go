package app

import (
	"github.com/gin-gonic/gin"

	HealthCheckInterface "misteraladin.com/jasmine/go-boilerplate/app/health-check"

	HCHandler "misteraladin.com/jasmine/go-boilerplate/app/health-check/handler"
)

// Define your route here
// Register the route on main.go with usecase as the parameter

func HealthCheckHttpHandler(r *gin.Engine, us HealthCheckInterface.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}

	route := r.Group("/test")
	route.GET("/health-check", handler.Check)
}
