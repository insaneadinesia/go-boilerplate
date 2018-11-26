package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"misteraladin.com/jasmine/go-boilerplate/config"

	// Register Gorm Mysql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// Register Go Sql Driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	appConfig = config.Config.App
	dbConfig  = config.Config.DB
	mysqlConn *gorm.DB
	err       error
	err1      error
)

// initialize database
func init() {
	if dbConfig.Driver == "mysql" {
		setupMysqlConn()
	}
}

// setupMysqlConn: setup mysql database connection using the configuration from config.yml
func setupMysqlConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	mysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = mysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}

	if appConfig.ENV != "local" {
		mysqlConn.LogMode(false)
	} else {
		mysqlConn.LogMode(true)
	}
	// mysqlConn.DB().SetMaxIdleConns(mysql.MaxIdleConns)
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}
