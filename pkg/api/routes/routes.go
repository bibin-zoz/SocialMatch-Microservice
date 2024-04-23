package routes

import (
	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, adminhandler *handlers.AdminHandler) {
	router.GET("/ping", handlers.PingHandler)
	router.POST("/login", userHandler.Userlogin)
	router.GET("/verify", userHandler.UserOtpReq)
	router.POST("/verify", userHandler.UserOtpVerification)
	router.POST("/signup", userHandler.UserSignup)

	router.POST("/admin/login", adminhandler.LoginHandler)

	router.PUT("/profile", userHandler.UserEditDetails)

	router.GET("/admin/users", userHandler.GetAllUsers)
	router.PUT("/admin/users", adminhandler.UpdateUserStatus)
	router.GET("/admin/interests", adminhandler.GetInterests)              // Register the route for fetching interests
	router.GET("/admin/preferences", adminhandler.GetPreferences)          // Register the route for fetching preferences
	router.POST("/admin/interests", adminhandler.AddInterest)              // Register the route for adding interest
	router.PUT("/admin/interests", adminhandler.EditInterest)              // Register the route for editing interest
	router.DELETE("/admin/interests/:id", adminhandler.DeleteInterest)     // Register the route for deleting interest
	router.POST("/admin/preferences", adminhandler.AddPreference)          // Register the route for adding preference
	router.PUT("/admin/preferences", adminhandler.EditPreference)          // Register the route for editing preference
	router.DELETE("/admin/preferences/:id", adminhandler.DeletePreference) // Register the route for deleting preference

	// new routes
	router.GET("/user/:user_id/interests", userHandler.GetUserInterests)
	router.GET("/user/:user_id/preferences", userHandler.GetUserPreferences)
	router.POST("/user/interests", userHandler.AddUserInterest)
	router.PUT("/user/interests", userHandler.EditUserInterest)
	router.DELETE("/user/interests", userHandler.DeleteUserInterest)
	router.POST("/user/preferences", userHandler.AddUserPreference)
	router.PUT("/user/preferences", userHandler.EditUserPreference)
	router.DELETE("/user/preferences", userHandler.DeleteUserPreference)

}
