package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/helper"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoomHandler struct {
	GRPC_RoomClient interfaces.RoomClient
}

func NewRoomHandler(RoomClient interfaces.RoomClient) *RoomHandler {
	return &RoomHandler{
		GRPC_RoomClient: RoomClient,
	}
}

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

	c.JSON(http.StatusCreated, createdRoom)
}

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

	c.JSON(http.StatusOK, editedRoom)
}

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

	c.JSON(http.StatusOK, changedRoom)
}

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

	c.JSON(http.StatusOK, updatedRoom)
}

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
	c.JSON(http.StatusOK, joinRequests)
}
func (rh *RoomHandler) GetAllRooms(c *gin.Context) {
	rooms, err := rh.GRPC_RoomClient.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all rooms"})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

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

	c.JSON(http.StatusOK, members)
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

	c.JSON(http.StatusCreated, sentMessage)
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
