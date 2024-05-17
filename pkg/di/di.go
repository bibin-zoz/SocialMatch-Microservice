package di

import (
	server "github.com/bibin-zoz/social-match-userauth-svc/pkg"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/api/service"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/client"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/db"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/repository"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, mongoClient, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	interestClient := client.NewInterestClient(cfg)
	connectionClient := client.NewConnectionsClient(cfg)
	chatClient := client.NewChatClient(cfg)
	userRepository := repository.NewUserRepository(gormDB, mongoClient)
	// kafkaProducer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	// kafkaConsumer := kafka.NewConsumer(cfg.KafkaBrokers, cfg.KafkaTopic, "1")
	userUsecase := usecase.NewUserUseCase(userRepository, cfg, interestClient, connectionClient, chatClient)
	// go userUsecase.ConsumeAndProcessMessages()
	userServiceServer := service.NewUserServer(userUsecase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
