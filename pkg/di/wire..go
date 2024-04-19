package di

import (
	"fmt"

	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/bibin-zoz/api-gateway/pkg/client"
	"github.com/bibin-zoz/api-gateway/pkg/config"
	server "github.com/bibin-zoz/api-gateway/pkg/server"
)

// func InitializeApi() *server.Server {

// 	server := server.NewServer()
// 	return server

// }
func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	fmt.Println("initializing")
	userClient := client.NewUserClient(cfg)
	userHandler := handlers.NewUserHandler(userClient)
	adminClient := client.NewAdminClient(cfg)
	adminHandler := handlers.NewAdminHandler(adminClient, userClient)

	serverHTTP := server.NewServerHTTP(userHandler, adminHandler)

	return serverHTTP, nil
}
