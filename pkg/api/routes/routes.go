package routes

import (
	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/bibin-zoz/api-gateway/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, adminhandler *handlers.AdminHandler) {
	router.GET("/ping", handlers.PingHandler)
	router.POST("/login", userHandler.Userlogin)
	router.GET("/verify", userHandler.UserOtpReq)
	router.POST("/verify", userHandler.UserOtpVerification)
	router.POST("/signup", userHandler.UserSignup)

	router.POST("/admin/login", adminhandler.LoginHandler)

	router.Use(middleware.UserAuthMiddleware())
	{
		router.PUT("/profile", userHandler.UserEditDetails)

	}

	router.Use(middleware.AdminAuthMiddleware())
	{
		router.GET("/admin/users", userHandler.GetAllUsers)
	}

	// Add more routes here

}
