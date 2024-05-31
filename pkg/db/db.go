package db

import (
	"fmt"
	"log"

	"github.com/bibin-zoz/social-match-admin-svc/pkg/config"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	dbName := cfg.DBName
	var exists bool

	// Check if the database exists
	err := db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		log.Fatal(err)
	}

	// Create the database if it does not exist
	if !exists {
		err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			log.Fatal(err)
		}
		log.Println("created database " + dbName)
	}

	// Connect to the newly created database
	db, err = gorm.Open(postgres.Open(psqlInfo+" dbname="+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate the domain models
	err = db.AutoMigrate(&domain.Admin{}, &domain.Interest{}, &domain.Preference{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
