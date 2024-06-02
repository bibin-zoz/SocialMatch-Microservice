package db

import (
	"context"
	"fmt"
	"log"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase establishes connections to both PostgreSQL and MongoDB databases
func ConnectDatabase(cfg config.Config) (*gorm.DB, *mongo.Database, error) {
	// PostgreSQL connection
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		return nil, nil, dbErr
	}

	dbName := cfg.DBName
	var exists bool

	// Check if the database exists
	err := db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		return nil, nil, err
	}

	// Create the database if it does not exist
	if !exists {
		err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			return nil, nil, err
		}
		log.Println("Created database " + dbName)
	}

	// Connect to the newly created database
	db, err = gorm.Open(postgres.Open(psqlInfo+" dbname="+dbName), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	// MongoDB connection
	mongoURI := fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBHost, cfg.MongoDBPort)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, nil, err
	}

	// Check the MongoDB connection
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}
	mongoDB := mongoClient.Database(cfg.MongoDBDatabase)

	// AutoMigrate the domain models
	err = db.AutoMigrate(&domain.User{}, &domain.UserInterest{}, &domain.UserPreference{}, &domain.Connections{}, &domain.UserMessage{}, &domain.Media{})
	if err != nil {
		return nil, nil, err
	}

	return db, mongoDB, nil
}
