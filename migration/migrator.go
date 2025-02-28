package migration

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"
	"time"

	"gorm.io/gorm"
)

var migrator = &Migrator{
	Migrations: map[string]*Migration{},
}

// Init : initialize schema migrations
func Init(db *gorm.DB) (*Migrator, error) {
	migrator.db = db

	// Create schema_migrations table to remember which migrations were executed.
	if err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		batch INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
		`).Error; err != nil {
		fmt.Println("Unable to create schema_migrations table", err)
		return migrator, err
	}

	rows, err := db.Raw("SELECT * FROM schema_migrations;").Rows()
	if err != nil {
		return migrator, err
	}
	defer rows.Close()

	for rows.Next() {
		var mg Migration
		err := db.ScanRows(rows, &mg)
		if err != nil {
			return migrator, err
		}

		if migrator.Migrations[mg.Name] != nil {
			migrator.Migrations[mg.Name].ID = mg.ID
			migrator.Migrations[mg.Name].Name = mg.Name
			migrator.Migrations[mg.Name].Batch = mg.Batch
			migrator.Migrations[mg.Name].done = true
		}

		if migrator.MaxBatch < mg.Batch {
			migrator.MaxBatch = mg.Batch
		}
	}

	return migrator, nil
}

func Create(name string) error {
	version := time.Now().Format("20060102150405")

	in := struct {
		Version string
		Name    string
	}{
		Version: version,
		Name:    name,
	}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles("./migration/template.txt"))
	if err := t.Execute(&out, in); err != nil {
		return errors.New("Unable to execute template: " + err.Error())
	}

	f, err := os.Create(fmt.Sprintf("./migration/%s_%s.go", version, name))

	if err != nil {
		return errors.New("Unable to create migration file: " + err.Error())
	}

	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		return errors.New("Unable to write to migration file: " + err.Error())
	}

	fmt.Println("Generated new migration file...", f.Name())
	return nil
}

// AddMigration : add new migration version
func (m *Migrator) AddMigration(mg *Migration) {
	// Add the migration to the hash with version as key
	m.Migrations[mg.Name] = mg
}

// Up .
func (m *Migrator) Up() error {
	tx := m.db.Begin()

	for _, mg := range m.Migrations {
		if mg.done {
			continue
		}

		fmt.Println("Running migration", mg.Name)
		if err := mg.Up(tx); err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Exec("INSERT INTO schema_migrations (name, batch) VALUES($1, $2)", mg.Name, m.MaxBatch+1).Error; err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println("Finished running migration", mg.Name)
	}

	tx.Commit()

	return nil
}

// Down .
func (m *Migrator) Down() error {
	tx := m.db.Begin()

	for _, mg := range m.Migrations {
		if !mg.done || mg.Batch != m.MaxBatch {
			continue
		}

		fmt.Println("Reverting migration", mg.Name)
		if err := mg.Down(tx); err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Exec("DELETE FROM schema_migrations WHERE id = $1", mg.ID).Error; err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println("Finished reverting migration", mg.Name)
	}

	tx.Commit()

	return nil
}

// MigrationStatus .
func (m *Migrator) MigrationStatus() error {
	for _, mg := range m.Migrations {
		if mg.done {
			fmt.Printf("%s", fmt.Sprintf("Migration %s... completed", mg.Name))
		} else {
			fmt.Printf("%s", fmt.Sprintf("Migration %s... pending", mg.Name))
		}
	}

	return nil
}
