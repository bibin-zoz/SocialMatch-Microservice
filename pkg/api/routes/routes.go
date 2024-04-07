package routes

import (
	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	router.GET("/ping", handlers.PingHandler)
	router.POST("/login", userHandler.Userlogin)
	router.POST("/signup", userHandler.UserSignup)

	// Add more routes here

}
