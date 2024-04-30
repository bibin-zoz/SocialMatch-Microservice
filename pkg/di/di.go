package di

import (
	service "github.com/bibin-zoz/social-match-admin-svc/pkg/api/service"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/config"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/db"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/repository"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/server"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	interestUseCase := usecase.NewInterestUseCase()

	adminServiceServer := service.NewAdminServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
