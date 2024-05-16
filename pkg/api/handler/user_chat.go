package handlers

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/helper"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// UserChatHandler manages WebSocket connections for user chats
type UserChatHandler struct {
	GRPC_Client   interfaces.UserClient
	Upgrader      websocket.Upgrader
	Room          map[string]*models.WebrtcRoom
	ConnectionsMu sync.Mutex
}

// NewUserChatHandler creates a new instance of UserChatHandler
func NewUserChatHandler(UserClient interfaces.UserClient) *UserChatHandler {
	return &UserChatHandler{
		Upgrader: websocket.Upgrader{},
		Room:     make(map[string]*models.WebrtcRoom),
	}
}

func (uch *UserChatHandler) HandleWebSocket(c *gin.Context) {
	ws, err := uch.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer ws.Close()

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		return
	}
	receiverID, err := strconv.ParseInt(c.Query("receiver_id"), 10, 64)
	if err != nil {
		log.Printf("Invalid receiver ID: %v", err)
		return
	}

	roomID := helper.GenerateRoomID(userID, receiverID)
	uch.ConnectionsMu.Lock()
	if _, ok := uch.Room[roomID]; !ok {
		uch.Room[roomID] = &models.WebrtcRoom{
			User1:       userID,
			User2:       receiverID,
			Connections: []*websocket.Conn{ws},
			Ch:          make(chan *models.WebrtcMessage),
		}
		go helper.BroadcastMessages(roomID, uch.Room, &uch.ConnectionsMu)
	} else {
		uch.Room[roomID].Connections = append(uch.Room[roomID].Connections, ws)
	}
	uch.ConnectionsMu.Unlock()
	defer func() {
		uch.ConnectionsMu.Lock()
		for i, conn := range uch.Room[roomID].Connections {
			if conn == ws {
				uch.Room[roomID].Connections = append(uch.Room[roomID].Connections[:i], uch.Room[roomID].Connections[i+1:]...)
				break
			}
		}
		uch.ConnectionsMu.Unlock()
	}()

	for {
		var msg models.WebrtcMessage
		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		// Save message to database
		// msg.Time = time.Now()
		userMessage := models.UserMessage{
			SenderID:   uint(msg.SenderID),
			RecipentID: uint(msg.ReceiverID),
			Content:    msg.Message,
			CreatedAt:  time.Now(),
		}
		if err := helper.SendMessageKafka(userMessage, c); err != nil {
			log.Printf("Error saving message: %v", err)
			continue
		}
		uch.Room[roomID].Ch <- &msg
	}
}
