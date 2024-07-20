package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/helper"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

type RoomHandler struct {
	GRPC_RoomClient interfaces.RoomClient
	Upgrader        websocket.Upgrader
	rooms           map[string]*models.GroupChatRoomRoom
	connectionsMu   sync.Mutex
	// Room            models.GroupChatRoomRoom
}

func NewRoomHandler(RoomClient interfaces.RoomClient) *RoomHandler {
	return &RoomHandler{
		GRPC_RoomClient: RoomClient,
		Upgrader:        websocket.Upgrader{},
		rooms:           make(map[string]*models.GroupChatRoomRoom),
		// Room:            models.GroupChatRoomRoom{},
	}
}

// CreateRoom godoc
// @Summary Create a new room
// @Description Create a new room for group chat
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Param roomData body models.RoomData true "Room Data"
// @Success 201 {object} models.Room
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /rooms [post]
func (rh *RoomHandler) CreateRoom(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)

	// Extract user ID from token
	userID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	fmt.Println("id", userID)

	var roomData models.RoomData
	if err := c.ShouldBindJSON(&roomData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}
	err = validator.New().Struct(roomData)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	createdRoom, err := rh.GRPC_RoomClient.CreateRoom(roomData, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Room created successfully", createdRoom, nil)

	c.JSON(http.StatusCreated, success)
}

// @Summary Edit a room
// @Description Edit an existing room
// @Tags Rooms
// @Accept json
// @Produce json
// @Param roomData body models.RoomData true "Room Data"
// @Success 200 {object} models.Room
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /rooms [put]
func (rh *RoomHandler) EditRoom(c *gin.Context) {
	var roomData models.RoomData
	if err := c.ShouldBindJSON(&roomData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	editedRoom, err := rh.GRPC_RoomClient.EditRoom(roomData)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to edit room: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	success := response.ClientResponse(http.StatusCreated, "room edited successfully", editedRoom, nil)

	c.JSON(http.StatusOK, success)
}

// @Summary Change room status
// @Description Change the status of a room
// @Tags Rooms
// @Accept json
// @Produce json
// @Param requestData body object true "Change room status request"
// @Property requestData.room_id uint32 required "Room ID"
// @Property requestData.status string required "New status of the room"
// @Success 200 {object} models.Room
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /room [patch]
func (rh *RoomHandler) ChangeRoomStatus(c *gin.Context) {
	var requestData struct {
		RoomID uint32 `json:"room_id"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	changedRoom, err := rh.GRPC_RoomClient.ChangeRoomStatus(requestData.RoomID, requestData.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change room status"})
		return
	}
	success := response.ClientResponse(http.StatusCreated, "room status changed successfully", changedRoom, nil)

	c.JSON(http.StatusOK, success)
}

// @Summary Add members to a room
// @Description Add members to an existing room
// @Tags Rooms
// @Accept json
// @Produce json
// @Param requestData body object true "Add members request"
// @Property requestData.room_id uint32 required "Room ID"
// @Property requestData.user_ids array of uint32 required "User IDs to add"
// @Success 200 {object} models.Room
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /room/members [put]
func (rh *RoomHandler) AddMembersToRoom(c *gin.Context) {
	var requestData struct {
		RoomID  uint32   `json:"room_id"`
		UserIDs []uint32 `json:"user_ids"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	updatedRoom, err := rh.GRPC_RoomClient.AddMembersToRoom(requestData.RoomID, requestData.UserIDs)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to add users: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	success := response.ClientResponse(http.StatusCreated, "member added to room successfully", updatedRoom, nil)

	c.JSON(http.StatusOK, success)
}

// @Summary Get room join requests
// @Description Get a list of join requests for a room
// @Tags Rooms
// @Produce json
// @Param requestData body object true "Get join requests request"
// @Property requestData.room_id uint32 required "Room ID"
// @Success 200 {array} models.Room
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /rooms/members/requests [get]
func (rh *RoomHandler) GetRoomJoinRequests(c *gin.Context) {
	var requestData struct {
		RoomID uint32 `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	joinRequests, err := rh.GRPC_RoomClient.GetRoomJoinRequests(requestData.RoomID)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get join req : %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	success := response.ClientResponse(http.StatusCreated, "requested sented successfully", joinRequests, nil)

	c.JSON(http.StatusOK, success)
}

// GetAllRooms godoc
// @Summary Get all rooms
// @Description Get a list of all rooms
// @Tags Rooms
// @Produce  json
// @Success 200 {array} []models.Room
// @Failure 500 {object} response.Response
// @Router /user/room [get]
func (rh *RoomHandler) GetAllRooms(c *gin.Context) {
	rooms, err := rh.GRPC_RoomClient.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all rooms"})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

// @Summary Get all rooms
// @Description Get a list of all rooms
// @Tags Rooms
// @Produce json
// @Success 200 {array} []models.Room
// @Failure 500 {object} response.Response
// @Router /user/room [get]
func (rh *RoomHandler) GetRoomMembers(c *gin.Context) {
	fmt.Println("hii")
	roomID, err := strconv.ParseUint(c.Param("room_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	members, err := rh.GRPC_RoomClient.GetRoomMembers(uint32(roomID))
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get room users: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	success := response.ClientResponse(http.StatusCreated, "members fetcehd successfully", members, nil)

	c.JSON(http.StatusOK, success)
}

func (rh *RoomHandler) SendMessage(c *gin.Context) {
	var requestData struct {
		Content string         `json:"content"`
		Media   []models.Media `json:"media"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	roomID, err := strconv.ParseUint(c.Param("room_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)
	senderID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Prepare the message object
	message := models.Message{
		RoomID:  uint(roomID),
		UserID:  uint(senderID),
		Content: requestData.Content,
		Media:   requestData.Media,
	}

	// Send the message
	sentMessage, err := rh.GRPC_RoomClient.SendMessage(message)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to send message: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	success := response.ClientResponse(http.StatusCreated, "msg sented successfully", sentMessage, nil)

	c.JSON(http.StatusCreated, success)
}

func (rh *RoomHandler) ReadMessages(c *gin.Context) {
	// Extract room ID from request context
	roomID, err := strconv.ParseUint(c.Param("room_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	messages, err := rh.GRPC_RoomClient.ReadMessages(uint32(roomID))
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to read messages: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, messages)
}
func (rh *RoomHandler) HandleWebSocket(c *gin.Context) {

	// Upgrade initial GET request to a WebSocket
	ws, err := rh.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer ws.Close()

	// Get user ID and group ID from query parameters
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		fmt.Println("userID", userID)
		log.Printf("Invalid user ID: %v", err)
		return
	}

	groupID, err := strconv.ParseInt(c.Query("group_id"), 10, 64)
	if err != nil {
		log.Printf("Invalid group ID: %v", err)
		return
	}

	roomID := strconv.FormatInt(groupID, 10) // Use group ID as room ID for simplicity

	// Create a new room if it doesn't exist
	rh.connectionsMu.Lock()
	if _, ok := rh.rooms[roomID]; !ok {
		rh.rooms[roomID] = &models.GroupChatRoomRoom{
			GroupID:     groupID,
			Connections: []*websocket.Conn{ws},
			Ch:          make(chan *models.UserMessage),
		}
		go broadcastMessages(roomID, rh.rooms)
	} else {
		rh.rooms[roomID].Connections = append(rh.rooms[roomID].Connections, ws)
	}
	rh.connectionsMu.Unlock()

	// Fetch previous messages from the database (if needed)

	// Remove connection when this function returns
	defer func() {
		rh.connectionsMu.Lock()
		for i, conn := range rh.rooms[roomID].Connections {
			if conn == ws {
				rh.rooms[roomID].Connections = append(rh.rooms[roomID].Connections[:i], rh.rooms[roomID].Connections[i+1:]...)
				break
			}
		}
		rh.connectionsMu.Unlock()
	}()

	// Listen for incoming messages
	for {
		var msg models.UserMessage
		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		// Save message to database
		// msg.Time = time.Now()
		// if err := saveMessage(&msg); err != nil {
		// 	log.Printf("Error saving message: %v", err)
		// 	continue
		// }

		// Broadcast message to all members in the room
		rh.rooms[roomID].Ch <- &msg
	}
}

func broadcastMessages(roomID string, rooms map[string]*models.GroupChatRoomRoom) {
	room := rooms[roomID]
	for msg := range room.Ch {
		for _, conn := range room.Connections {
			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("Error sending message to connection: %v", err)
			}
		}
	}
}
