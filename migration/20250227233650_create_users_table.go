package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250227233650_create_users_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250227233650_create_users_table_up,
		Down:      mig_20250227233650_create_users_table_down,
	})
}

func mig_20250227233650_create_users_table_up(tx *gorm.DB) error {
	err := tx.Exec(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS users (
			uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			name VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL UNIQUE,
			sub_district_id INTEGER NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP
		);`).Error

	return err
}

func mig_20250227233650_create_users_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS users;`).Error
	return err
}
