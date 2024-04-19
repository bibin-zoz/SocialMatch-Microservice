package di

import (
	server "github.com/bibin-zoz/social-match-userauth-svc/pkg"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/api/service"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/db"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/repository"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(gormDB)
	adminUseCase := usecase.NewUserUseCase(userRepository, cfg)

	userServiceServer := service.NewUserServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
