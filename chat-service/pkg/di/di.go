package di

import (
	server "github.com/bibin-zoz/social-match-chat-svc/pkg"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/api/service"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/config"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/db"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/repository"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, mongoClient, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewConnectionRepository(gormDB, mongoClient)
	// kafkaProducer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	// kafkaConsumer := kafka.NewConsumer(cfg.KafkaBrokers, cfg.KafkaTopic, "1")

	userUsecase := usecase.NewChatUseCase(userRepository, cfg)
	go userUsecase.ConsumeAndProcessMessages()
	userServiceServer := service.NewChatServer(userUsecase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
