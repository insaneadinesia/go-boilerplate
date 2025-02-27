package health_check

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/internal/app/repository"
)

type HealthCheckUsecase interface {
	HealthCheck(ctx context.Context) (resp StatusCheck, err error)
}

type usecase struct {
	healthCheckRepository repository.HealthCheck
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (u *usecase) SetHealthCheckRepository(repo repository.HealthCheck) *usecase {
	u.healthCheckRepository = repo
	return u
}

func (u *usecase) Validate() HealthCheckUsecase {
	if u.healthCheckRepository == nil {
		panic("healthCheckRepository is nil")
	}

	return u
}
