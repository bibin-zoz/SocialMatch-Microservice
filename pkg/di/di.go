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
	userRepository := repository.NewUserRepository(gormDB, mongoClient)
	adminUseCase := usecase.NewUserUseCase(userRepository, cfg, interestClient)

	userServiceServer := service.NewUserServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
