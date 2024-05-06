package db

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase establishes connections to both PostgreSQL and MongoDB databases
func ConnectDatabase(cfg config.Config) (*gorm.DB, *mongo.Client, error) {
	// PostgreSQL connection
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// MongoDB connection
	mongoURI := fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBHost, cfg.MongoDBPort)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}
	db.AutoMigrate(&domain.User{}, &domain.UserInterest{}, &domain.UserPreference{})

	return db, mongoClient, dbErr
}
