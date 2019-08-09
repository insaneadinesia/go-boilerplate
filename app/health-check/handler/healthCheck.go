package handler

import (
	"time"

	"misteraladin.com/jasmine/go-boilerplate/lang"

	"github.com/gin-gonic/gin"
	Base "misteraladin.com/jasmine/go-boilerplate/app/api/handler"
	HealthCheckInterface "misteraladin.com/jasmine/go-boilerplate/app/health-check"
)

type HealthCheckResponse struct {
	HealthStatus string    `json:"health_status"`
	DBTimestamp  time.Time `json:"database_timestamp"`
}

type HealthCheckHandler struct {
	HealthCheckUsecase HealthCheckInterface.IHealthCheckUsecase
}

// Handler just handle how data come and how data will serve
// To process the incoming data, you need to connect to the usecase via interface

func (a *HealthCheckHandler) Check(c *gin.Context) {
	healthCheck := a.HealthCheckUsecase.GetDBTimestamp()
	res := &HealthCheckResponse{
		HealthStatus: lang.Translate("health_status", nil),
		DBTimestamp:  healthCheck.CurrentTimestamp,
	}

	Base.RespondJSON(c, res)
	return
}
