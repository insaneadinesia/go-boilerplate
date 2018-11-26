package healthcheck

import "misteraladin.com/jasmine/go-boiler-plate/models"

type IHealthCheckRepository interface {
	GetDBTimestamp() models.HealthCheck
}
