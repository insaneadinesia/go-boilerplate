package usecase

import (
	HealthCheckInterface "misteraladin.com/jasmine/go-boiler-plate/app/health-check"
	"misteraladin.com/jasmine/go-boiler-plate/models"
)

// Define your usecase struct
// You can make the struct with repository or another usecase
// To access the repository or the usecase you need an interface
type HealthCheckUsecase struct {
	HealthCheckRepository HealthCheckInterface.IHealthCheckRepository
}

// Define an exported function. Call this in main.go and bind your parameters to the struct
func NewHealthCheckUsecase(h HealthCheckInterface.IHealthCheckRepository) HealthCheckInterface.IHealthCheckUsecase {
	return &HealthCheckUsecase{
		HealthCheckRepository: h,
	}
}

// Define your custom functions
// You can put any logic on the usecase
// And to connect to the database, you can use the repository via interface
// To export the functions, you need to register function to the interface

func (a *HealthCheckUsecase) GetDBTimestamp() models.HealthCheck {
	healthCheck := a.HealthCheckRepository.GetDBTimestamp()
	return healthCheck
}
