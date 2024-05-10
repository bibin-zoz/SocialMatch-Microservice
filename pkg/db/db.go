package db

import (
	"fmt"

	"github.com/bibin-zoz/social-match-room-svc/pkg/config"
	"github.com/bibin-zoz/social-match-room-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Room{}, &domain.RoomMember{}, &domain.RoomJoinReq{}, &domain.Message{})
	return db, dbErr

}
