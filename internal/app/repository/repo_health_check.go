package repository

import (
	"context"

	"gorm.io/gorm"
)

type HealthCheck interface {
	PingDB(ctx context.Context) (err error)
}

type heatlhCheck struct {
	db *gorm.DB
}

func NewHealthCheckRepository(db *gorm.DB) HealthCheck {
	if db == nil {
		panic("database is nil")
	}

	return &heatlhCheck{
		db: db,
	}
}

func (r *heatlhCheck) PingDB(ctx context.Context) (err error) {
	db, _ := r.db.DB()
	err = db.PingContext(ctx)
	return
}
