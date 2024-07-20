package routes

import (
	handlers "github.com/bibin-zoz/api-gateway/pkg/api/handler"
	"github.com/bibin-zoz/api-gateway/pkg/api/middleware"
	"github.com/bibin-zoz/api-gateway/pkg/api/websocketserver"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, userAuthHandler *handlers.UserAuthHandler, adminhandler *handlers.AdminHandler, roomHandler *handlers.RoomHandler, userChatHandler *handlers.UserChatHandler, videocallHandler *handlers.VideoCallHandler) {
	// router.Use(middleware.CORSMiddleware())
	router.GET("/ping", handlers.PingHandler)
	router.POST("/login", userAuthHandler.Userlogin)
	router.GET("/verify", userAuthHandler.UserOtpReq)
	router.POST("/verify", userAuthHandler.UserOtpVerification)
	router.POST("/signup", userAuthHandler.UserSignup)

	router.POST("/admin/login", adminhandler.LoginHandler)

	router.PUT("/user/profile", userHandler.UserEditDetails)
	router.PATCH("/user/profile", userHandler.UpdateProfilePhoto)

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

	//rooms
	router.GET("/user/room", middleware.UserAuthMiddleware(), roomHandler.GetAllRooms)
	router.POST("/user/room", middleware.UserAuthMiddleware(), roomHandler.CreateRoom)
	router.PUT("/user/room", middleware.UserAuthMiddleware(), roomHandler.EditRoom)
	router.PUT("/user/room", middleware.UserAuthMiddleware(), roomHandler.ChangeRoomStatus)
	router.POST("/user/room/members", middleware.UserAuthMiddleware(), roomHandler.AddMembersToRoom)
	router.GET("/user/room/members/:room_id", roomHandler.GetRoomMembers)
	router.GET("/user/room/members/requests", roomHandler.GetRoomJoinRequests)
	// hub := Hub.NewHub()
	router.GET("/wsroom", roomHandler.HandleWebSocket)

	router.POST("/user/room/:room_id", roomHandler.SendMessage)
	router.GET("/user/room/:room_id", roomHandler.ReadMessages)
	// router.POST("/user/message", userHandler.SendMessageHandler)
	// router.POST("/user/message", userHandler.SendMessageKafka)

	//friend
	router.GET("/user/connections", userHandler.GetConnections)
	router.POST("/user/connections", userHandler.UserFollow)
	router.DELETE("/user/connections", userHandler.BlockUser)

	router.GET("ws", userChatHandler.HandleWebSocket)
	router.GET("/user/message", userHandler.ReadMessages)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")

	router.GET("/exit", videocallHandler.ExitPage)
	router.GET("/error", videocallHandler.ErrorPage)
	router.GET("/videocall", videocallHandler.IndexedPage)
	router.GET("/wsvideocall", videocallHandler.HandleWebSocket)

	router.POST("/create", websocketserver.CreateRoomRequestHandler)
	router.GET("/join", websocketserver.JoinRoomRequestHandler)
}
