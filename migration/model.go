package migration

import (
	"time"

	"gorm.io/gorm"
)

// Migration information & command struct
type Migration struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Batch     int       `db:"batch"`
	CreatedAt time.Time `db:"created_at"`

	Up   func(*gorm.DB) error `gorm:"-"`
	Down func(*gorm.DB) error `gorm:"-"`

	done bool `gorm:"-"`
}

// Migrator : collection of migration
type Migrator struct {
	db         *gorm.DB
	Migrations map[string]*Migration
	MaxBatch   int
}
