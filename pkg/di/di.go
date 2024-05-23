package di

import (
	"github.com/bibin-zoz/social-match-admin-svc/pkg/api/service"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/config"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/db"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/repository"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/server"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	// Connect to the database
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	// Create repository instances
	adminRepository := repository.NewAdminRepository(gormDB)
	interestRepository := repository.NewInterestRepository(gormDB)

	// Create use case instances
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	interestUseCase := usecase.NewInterestUseCase(interestRepository)

	// Create service instances
	adminService := service.NewAdminServer(adminUseCase)
	interestService := service.NewInterestServiceServer(interestUseCase)

	// Create and start the gRPC server
	grpcServer, err := server.NewGRPCServer(cfg, adminService, interestService)
	if err != nil {
		return nil, err
	}

	return grpcServer, nil
}
