package di

import (
	server "github.com/bibin-zoz/social-match-room-svc/pkg"
	"github.com/bibin-zoz/social-match-room-svc/pkg/api/service"
	"github.com/bibin-zoz/social-match-room-svc/pkg/config"
	"github.com/bibin-zoz/social-match-room-svc/pkg/db"
	"github.com/bibin-zoz/social-match-room-svc/pkg/repository"
	"github.com/bibin-zoz/social-match-room-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	roomRepository := repository.NewRoomRepository(gormDB)
	roomUseCase := usecase.NewRoomUseCase(roomRepository, cfg)

	roomServiceServer := service.NewRoomServer(roomUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, roomServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
