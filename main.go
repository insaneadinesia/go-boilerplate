package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"misteraladin.com/jasmine/go-boilerplate/config"
	gorm "misteraladin.com/jasmine/go-boilerplate/db"
	"misteraladin.com/jasmine/go-boilerplate/lang"
	"misteraladin.com/jasmine/go-boilerplate/redis"

	routes "misteraladin.com/jasmine/go-boilerplate/app"

	HCRepository "misteraladin.com/jasmine/go-boilerplate/app/health-check/repository"
	RDRepository "misteraladin.com/jasmine/go-boilerplate/app/redis/repository"

	HCUsecase "misteraladin.com/jasmine/go-boilerplate/app/health-check/usecase"
)

var appConfig = config.Config.App

func main() {
	r := gin.New()

	db := gorm.MysqlConn()
	redis := redis.RedisClient()

	r.Use(gin.Recovery())
	r.Use(localeMiddleware())

	// Register your repository here to the spesific connection which you'll use.
	// It can be database connection, redis connection, etc.
	// Delete redis if you not use it
	hcr := HCRepository.NewHealthCheckRepository(db)
	_ = RDRepository.NewRedisRepository(redis)

	// Register your usecase here. And add the parameter if any.
	hcu := HCUsecase.NewHealthCheckUsecase(hcr)

	// Register your usecase to the router here.
	routes.HealthCheckHttpHandler(r, hcu)

	// Run application
	if err := r.Run(fmt.Sprintf(":%s", appConfig.HTTPPort)); err != nil {
		log.Fatal(err)
	}
}

func localeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch strings.ToLower(c.GetHeader("Accept-Language")) {
		case "id":
			config.Config.App.Locale = "id"
			break
		default:
			config.Config.App.Locale = "en"
			break
		}

		lang.LoadLanguage()
		c.Next()
	}
}
