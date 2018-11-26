package healthcheck

import "misteraladin.com/jasmine/go-boiler-plate/models"

type IHealthCheckUsecase interface {
	GetDBTimestamp() models.HealthCheck
}
