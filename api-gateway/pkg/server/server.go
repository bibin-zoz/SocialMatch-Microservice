package server

import (
	"log"

	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/bibin-zoz/api-gateway/pkg/api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandlers *handlers.UserHandler, userAuthHandler *handlers.UserAuthHandler, adminHandlers *handlers.AdminHandler, roomHandlers *handlers.RoomHandler, userChatHandlers *handlers.UserChatHandler, videocallHandler *handlers.VideoCallHandler) *ServerHTTP {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // Allow all methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},          // Allow specified headers
		ExposeHeaders:    []string{"Content-Length"},                                   // Expose Content-Length header
		AllowCredentials: true,
	}))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Swagger route
	routes.SetupRoutes(router, userHandlers, userAuthHandler, adminHandlers, roomHandlers, userChatHandlers, videocallHandler)

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
