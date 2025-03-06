package repository

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/internal/pkg/helper"
	"github.com/insaneadinesia/gobang/gotel"
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
	ctx, span := gotel.Otel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	db, _ := r.db.DB()
	err = db.PingContext(ctx)
	return
}
