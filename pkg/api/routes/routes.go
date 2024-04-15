package routes

import (
	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, adminhandler *handlers.AdminHandler) {
	router.GET("/ping", handlers.PingHandler)
	router.POST("/login", userHandler.Userlogin)
	router.GET("verification", userHandler.UserOtpReq)
	router.POST("verification", userHandler.UserOtpVerification)
	router.POST("/signup", userHandler.UserSignup)
	router.PUT("/profile", userHandler.UserEditDetails)

	router.POST("/admin/login", adminhandler.LoginHandler)

	// Add more routes here

}
