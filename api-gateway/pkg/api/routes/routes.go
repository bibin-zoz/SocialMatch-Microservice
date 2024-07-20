package routes

import (
	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/bibin-zoz/api-gateway/pkg/api/middleware"
	"github.com/bibin-zoz/api-gateway/pkg/api/websocketserver"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, userAuthHandler *handlers.UserAuthHandler, adminhandler *handlers.AdminHandler, roomHandler *handlers.RoomHandler, userChatHandler *handlers.UserChatHandler, videocallHandler *handlers.VideoCallHandler) {
	userAuth := router.Group("/auth")
	{
		userAuth.POST("/login", userAuthHandler.Userlogin)
		userAuth.GET("/verify", userAuthHandler.UserOtpReq)
		userAuth.POST("/verify", userAuthHandler.UserOtpVerification)
		userAuth.POST("/signup", userAuthHandler.UserSignup)
	}

	// Admin Authentication Routes
	adminAuth := router.Group("/admin")
	{
		adminAuth.POST("/login", adminhandler.LoginHandler)
	}

	// User Profile Routes
	userProfile := router.Group("/user/profile")
	{
		userProfile.PUT("", userHandler.UserEditDetails)
		userProfile.PATCH("", userHandler.UpdateProfilePhoto)
	}

	// Admin User Management Routes
	adminUser := router.Group("/admin/users")
	{
		adminUser.GET("", userHandler.GetAllUsers)
		adminUser.PUT("", adminhandler.UpdateUserStatus)
	}

	// User Interest and Preference Routes
	userData := router.Group("/user")
	{
		userData.GET("/:user_id/interests", userHandler.GetUserInterests)
		userData.GET("/:user_id/preferences", userHandler.GetUserPreferences)
		userData.POST("/interests", userHandler.AddUserInterest)
		userData.PUT("/interests", userHandler.EditUserInterest)
		userData.DELETE("/interests", userHandler.DeleteUserInterest)
		userData.POST("/preferences", userHandler.AddUserPreference)
		userData.PUT("/preferences", userHandler.EditUserPreference)
		userData.DELETE("/preferences", userHandler.DeleteUserPreference)
	}

	// Room Routes
	room := router.Group("/user/room", middleware.UserAuthMiddleware())
	{
		room.GET("", roomHandler.GetAllRooms)
		room.POST("", roomHandler.CreateRoom)
		room.PUT("", roomHandler.EditRoom)
		room.PATCH("", roomHandler.ChangeRoomStatus)
		room.POST("/members", roomHandler.AddMembersToRoom)
		room.GET("/members/:room_id", roomHandler.GetRoomMembers)
		room.GET("/members/requests", roomHandler.GetRoomJoinRequests)
	}
	// Admin Interest and Preference Routes
	adminData := router.Group("/admin")
	{
		adminData.GET("/interests", adminhandler.GetInterests)
		adminData.GET("/preferences", adminhandler.GetPreferences)
		adminData.POST("/interests", adminhandler.AddInterest)
		adminData.PUT("/interests", adminhandler.EditInterest)
		adminData.DELETE("/interests/:id", adminhandler.DeleteInterest)
		adminData.POST("/preferences", adminhandler.AddPreference)
		adminData.PUT("/preferences", adminhandler.EditPreference)
		adminData.DELETE("/preferences/:id", adminhandler.DeletePreference)
	}

	// WebSocket Routes
	router.GET("/wsroom", roomHandler.HandleWebSocket)
	router.POST("/user/room/:room_id", roomHandler.SendMessage)
	router.GET("/user/room/:room_id", roomHandler.ReadMessages)

	// Friend Routes
	friend := router.Group("/user/connections")
	{
		friend.GET("", userHandler.GetConnections)
		friend.POST("", userHandler.UserFollow)
		friend.DELETE("", userHandler.BlockUser)
	}

	// Chat Routes
	chat := router.Group("/chat")
	{
		chat.GET("ws", userChatHandler.HandleWebSocket)
		chat.GET("/user/message", userHandler.ReadMessages)
	}

	// Video Call Routes
	videoCall := router.Group("/videocall")
	{
		videoCall.GET("/exit", videocallHandler.ExitPage)
		videoCall.GET("/error", videocallHandler.ErrorPage)
		videoCall.GET("/", videocallHandler.IndexedPage)
		videoCall.GET("/ws", videocallHandler.HandleWebSocket)
	}

	// WebSocket Server Routes
	webSocketServer := router.Group("/websocketserver")
	{
		webSocketServer.POST("/create", websocketserver.CreateRoomRequestHandler)
		webSocketServer.GET("/join", websocketserver.JoinRoomRequestHandler)
	}
	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")
}
