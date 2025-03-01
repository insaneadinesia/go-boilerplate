package container

import (
	"github.com/insaneadinesia/go-boilerplate/config"
	"github.com/insaneadinesia/go-boilerplate/internal/app/driver"
	"github.com/insaneadinesia/go-boilerplate/internal/app/repository"
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/health_check"
	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	"github.com/insaneadinesia/go-boilerplate/internal/app/wrapper/location_svc"
	"github.com/insaneadinesia/gobang/logger"
)

type Container struct {
	Config             config.Config
	HealthCheckUsecase health_check.HealthCheckUsecase
	UserUsecase        user.UserUsecase
}

func Setup() *Container {
	// Load Config
	cfg := config.Load()

	// Setup Driver
	db, _ := driver.NewPostgresDatabase(cfg)

	// Setup Tools
	logger.NewLogger(logger.Option{
		IsEnable:            cfg.LoggerEnable,
		EnableStackTrace:    cfg.LoggerEnableStackTrace,
		EnableMaskingFields: cfg.LoggerEnableMasking,
		MaskingFields:       cfg.LoggerMaskingFields,
	})

	// Setup Repository
	healthCheckRepository := repository.NewHealthCheckRepository(db)
	userRepository := repository.NewUserRepository(db)

	// Setup Wrapper
	locationSvcWrapper := location_svc.NewWrapper().SetConfig(cfg).Setup().Validate()

	// Setup Usecase
	healthCheckUsecase := health_check.NewUsecase().SetHealthCheckRepository(healthCheckRepository).Validate()
	userUsecase := user.NewUsecase().SetUserRepository(userRepository).SetLocationSvcWrapper(locationSvcWrapper).Validate()

	return &Container{
		Config:             cfg,
		HealthCheckUsecase: healthCheckUsecase,
		UserUsecase:        userUsecase,
	}
}
