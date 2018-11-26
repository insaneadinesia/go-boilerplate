package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"misteraladin.com/jasmine/go-boiler-plate/config"
	"misteraladin.com/jasmine/go-boiler-plate/db"
)

var appConfig = config.Config.App

func main() {
	r := gin.New()

	_ = gorm.MysqlConn()

	r.Use(gin.Recovery())

	if err := r.Run(fmt.Sprintf(":%s", appConfig.HTTPPort)); err != nil {
		log.Fatal(err)
	}
}
