package cmd

import (
	"log"
	"strings"

	"github.com/insaneadinesia/go-boilerplate/config"
	"github.com/insaneadinesia/go-boilerplate/internal/app/driver"
	"github.com/insaneadinesia/go-boilerplate/migration"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	cfg := config.Load()
	db, _ := driver.NewPostgresDatabase(cfg)

	return db
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migration file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Println("Unable to read flag `name`", err.Error())
			return
		}

		if err := migration.Create(name); err != nil {
			log.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// setup database connection
		db := initDB()

		migrator, err := migration.Init(db)
		if err != nil {
			log.Println("Unable to fetch migrator", err.Error())
			return
		}

		log.Println("Running migration...")
		err = migrator.Up()
		if err != nil {
			log.Println("Unable to run `up` migrations", err.Error())
			return
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// setup database connection
		db := initDB()

		migrator, err := migration.Init(db)
		if err != nil {
			log.Println("Unable to fetch migrator", err.Error())
			return
		}

		err = migrator.Down()
		if err != nil {
			log.Println("Unable to run `down` migrations")
			return
		}
	},
}

var migrateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()

		migrator, err := migration.Init(db)
		if err != nil {
			log.Println("Unable to fetch migrator")
			return
		}

		if err := migrator.MigrationStatus(); err != nil {
			log.Println("Unable to fetch migration status")
			return
		}
	},
}

func initMigrate() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				log.Fatalf("unknown command %q", strings.Join(args, " "))
			}

			_ = cmd.Help()
		},
	}

	// Add "--name" flag to "create" command
	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")
	// set flag "name" required
	_ = migrateCreateCmd.MarkFlagRequired("name")

	migrateCmd.AddCommand(migrateCreateCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateStatusCmd)

	return migrateCmd
}

func init() {
	rootCmd.AddCommand(initMigrate())
}
