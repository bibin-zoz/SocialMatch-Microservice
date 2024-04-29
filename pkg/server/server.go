package server

import (
	"log"

	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/bibin-zoz/api-gateway/pkg/api/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandlers *handlers.UserHandler, adminHandlers *handlers.AdminHandler, roomHandlers *handlers.RoomHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())
	routes.SetupRoutes(router, userHandlers, adminHandlers, roomHandlers)
	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
