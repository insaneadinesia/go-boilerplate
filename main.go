package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"misteraladin.com/jasmine/go-boiler-plate/config"
	"misteraladin.com/jasmine/go-boiler-plate/db"

	routes "misteraladin.com/jasmine/go-boiler-plate/app"

	HCRepository "misteraladin.com/jasmine/go-boiler-plate/app/health-check/repository"

	HCUsecase "misteraladin.com/jasmine/go-boiler-plate/app/health-check/usecase"
)

var appConfig = config.Config.App

func main() {
	r := gin.New()

	db := gorm.MysqlConn()

	r.Use(gin.Recovery())

	// Register your repository here to the spesific connection which you'll use.
	// It can be database connection, redis connection, etc.
	hcr := HCRepository.NewHealthCheckRepository(db)

	// Register your usecase here. And add the parameter if any.
	hcu := HCUsecase.NewHealthCheckUsecase(hcr)

	// Register your usecase to the router here.
	routes.HealthCheckHttpHandler(r, hcu)

	// Run application
	if err := r.Run(fmt.Sprintf(":%s", appConfig.HTTPPort)); err != nil {
		log.Fatal(err)
	}
}
