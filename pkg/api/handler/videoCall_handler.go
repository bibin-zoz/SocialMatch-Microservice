package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v2"
)

type VideoCallHandler struct {
	mu                sync.Mutex
	rooms             map[string]*Room
	Connection_Client interfaces.ConnectionClient
}

type Room struct {
	mu           sync.Mutex
	Participants map[*websocket.Conn]bool
}

func NewVideoCallHandler() *VideoCallHandler {
	return &VideoCallHandler{
		rooms: make(map[string]*Room),
	}
}

func (v *VideoCallHandler) ExitPage(c *gin.Context) {
	c.HTML(http.StatusOK, "exit.html", nil)
}

func (v *VideoCallHandler) ErrorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", nil)
}

func (v *VideoCallHandler) IndexedPage(c *gin.Context) {
	room := c.DefaultQuery("room", "")
	c.HTML(http.StatusOK, "index.html", gin.H{"room": room})
}

func (v *VideoCallHandler) SetupRoutes(group *gin.RouterGroup) {
	group.GET("/ws", v.handleWebSocket)
}

func (v *VideoCallHandler) handleWebSocket(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	userID := c.DefaultQuery("userID", "")
	receiverID := c.DefaultQuery("receiverID", "")

	// Convert user ID and receiver ID to integers
	uid, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("Invalid user ID:", err)
		return
	}
	rid, err := strconv.Atoi(receiverID)
	if err != nil {
		fmt.Println("Invalid receiver ID:", err)
		return
	}

	// Check user connection
	connection, err := v.Connection_Client.CheckUserConnection(int32(uid), int32(rid))
	if err != nil {
		fmt.Println("Error checking user connection:", err)
		return
	}

	// Check if users are connected
	if !connection {
		fmt.Println("Users are not connected")
		return
	}

	roomId := c.DefaultQuery("room", "")
	v.mu.Lock()
	room, exists := v.rooms[roomId]
	if !exists {
		room = &Room{
			Participants: make(map[*websocket.Conn]bool),
		}
		v.rooms[roomId] = room
	}
	v.mu.Unlock()

	room.mu.Lock()
	room.Participants[conn] = true
	room.mu.Unlock()

	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		fmt.Println("Error creating PeerConnection:", err)
		return
	}
	defer peerConnection.Close()

	// Handle WebRTC signaling here
}
