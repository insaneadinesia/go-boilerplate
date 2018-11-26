package healthcheck

import "misteraladin.com/jasmine/go-boilerplate/models"

type IHealthCheckUsecase interface {
	GetDBTimestamp() models.HealthCheck
}
