package db

import (
	"fmt"
	"log"

	"github.com/bibin-zoz/social-match-room-svc/pkg/config"
	"github.com/bibin-zoz/social-match-room-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	// PostgreSQL connection
	fmt.Println("connecting database...")
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		return nil, dbErr
	}

	dbName := cfg.DBName
	var exists bool

	// Check if the database exists
	err := db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		return nil, err
	}

	// Create the database if it does not exist
	if !exists {
		err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			return nil, err
		}
		log.Println("Created database " + dbName)
	}

	// Connect to the newly created database
	db, err = gorm.Open(postgres.Open(psqlInfo+" dbname="+dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate the domain models
	err = db.AutoMigrate(&domain.Room{}, &domain.RoomMember{}, &domain.RoomJoinReq{}, &domain.Message{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
