package driver

import (
	"fmt"

	"github.com/insaneadinesia/go-boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(cfg config.Config) (db *gorm.DB, err error) {
	fmt.Println("Try New Database ...")

	dsn := cfg.GetPostgresDSN()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if cfg.DBEnableDebug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConn)
	sqlDB.SetConnMaxLifetime(cfg.DBMaxConnLifetime)

	return
}
