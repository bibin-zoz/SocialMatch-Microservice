package handlers

import (
	"net/http"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	"github.com/gin-gonic/gin"
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
	var roomData models.RoomData
	if err := c.ShouldBindJSON(&roomData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	createdRoom, err := rh.GRPC_RoomClient.CreateRoom(roomData)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit room"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add members to room"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get room join requests"})
		return
	}

	c.JSON(http.StatusOK, joinRequests)
}
