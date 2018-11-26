package healthcheck

import "misteraladin.com/jasmine/go-boilerplate/models"

type IHealthCheckRepository interface {
	GetDBTimestamp() models.HealthCheck
}
