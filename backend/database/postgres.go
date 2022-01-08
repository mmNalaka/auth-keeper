package database

import (
	"log"

	"github.com/mmnalaka/auth-keeper/app/models"
	"github.com/mmnalaka/auth-keeper/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresInstance struct {
	Db *gorm.DB
}

var Postgres *PostgresInstance

func ConnectPostgres() {
	db, err := gorm.Open(postgres.Open(config.Cfg.Postgres.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected to Postgres database!")
	db.Logger = logger.Default.LogMode(logger.Info)

	if config.Cfg.Postgres.RunMigrations {
		log.Println("Running database migrations")
		if err := db.AutoMigrate(&models.User{}, models.Permission{}, models.RefreshToken{}, models.Client{}, models.LoginHistory{}, models.Role{}); err != nil {
			log.Fatal("Failed to run database migrations. \n", err)
		}
	}

	Postgres = &PostgresInstance{
		Db: db,
	}
}
