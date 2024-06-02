package db

import (
	"fmt"
	"log"

	"github.com/bibin-zoz/social-match-connections-svc/pkg/config"
	"github.com/bibin-zoz/social-match-connections-svc/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase establishes connections to both PostgreSQL and MongoDB databases
func ConnectDatabase(cfg config.Config) (*gorm.DB, *mongo.Client, error) {
	// PostgreSQL connection
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

	// // MongoDB connection
	// mongoURI := fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBHost, cfg.MongoDBPort)
	// mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	// if err != nil {
	// 	return nil, nil, err
	// }

	// // Check the MongoDB connection
	// err = mongoClient.Ping(context.Background(), nil)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// AutoMigrate the domain models
	err = db.AutoMigrate(&domain.User{}, &domain.UserInterest{}, &domain.UserPreference{}, &domain.Connections{}, &domain.UserMessage{}, &domain.Media{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil, nil
}
