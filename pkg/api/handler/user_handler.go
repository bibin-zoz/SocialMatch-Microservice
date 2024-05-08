package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/helper"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type UserHandler struct {
	GRPC_Client interfaces.UserClient
}

func NewUserHandler(UserClient interfaces.UserClient) *UserHandler {
	return &UserHandler{
		GRPC_Client: UserClient,
	}
}
func (ur *UserHandler) UserSignup(c *gin.Context) {
	var SignupDetail models.UserSignup
	if err := c.ShouldBindJSON(&SignupDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to bind Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(SignupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UsersSignUp(SignupDetail)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}
func (ur *UserHandler) Userlogin(c *gin.Context) {
	var UserLoginDetail models.UserLogin
	if err := c.ShouldBindJSON(&UserLoginDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}
	err := validator.New().Struct(UserLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UserLogin(UserLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully logged in with password", user, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) UserEditDetails(c *gin.Context) {
	var EditDetails models.UserUpdateDetails
	fmt.Println("edit", EditDetails)
	if err := c.ShouldBindJSON(&EditDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(EditDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UserEditDetails(EditDetails)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User Details edited successfully", user, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) UserOtpReq(c *gin.Context) {
	var EmailOtp models.UserVerificationRequest
	if err := c.ShouldBindJSON(&EmailOtp); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Email not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(EmailOtp)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UserOtpRequest(EmailOtp)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Otp sented successfully", user, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) UserOtpVerification(c *gin.Context) {
	var EmailOtp models.Otp
	if err := c.ShouldBindJSON(&EmailOtp); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Email not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(EmailOtp)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	_, err = ur.GRPC_Client.UserOtpVerificationReq(EmailOtp)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "otp verified successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}
func (ur *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := ur.GRPC_Client.GetAllUsers()
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Users fetched successfully", users, nil)
	c.JSON(http.StatusCreated, success)

}
func (ur *UserHandler) AddUserInterest(c *gin.Context) {
	var addUserInterest models.AddUserInterestRequest
	if err := c.ShouldBindJSON(&addUserInterest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(addUserInterest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.AddUserInterest(addUserInterest.UserID, addUserInterest.InterestID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User interest added successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) EditUserInterest(c *gin.Context) {
	var editUserInterest models.EditUserInterestRequest
	if err := c.ShouldBindJSON(&editUserInterest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(editUserInterest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.EditUserInterest(editUserInterest.UserID, editUserInterest.InterestID, editUserInterest.NewInterestName)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User interest edited successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) DeleteUserInterest(c *gin.Context) {
	var deleteUserInterest models.DeleteUserInterestRequest
	if err := c.ShouldBindJSON(&deleteUserInterest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(deleteUserInterest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.DeleteUserInterest(deleteUserInterest.UserID, deleteUserInterest.InterestID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User interest deleted successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}
func (ur *UserHandler) AddUserPreference(c *gin.Context) {
	var addUserPreference models.AddUserPreferenceRequest
	if err := c.ShouldBindJSON(&addUserPreference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(addUserPreference)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.AddUserPreference(addUserPreference.UserID, addUserPreference.PreferenceID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User preference added successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) EditUserPreference(c *gin.Context) {
	var editUserPreference models.EditUserPreferenceRequest
	if err := c.ShouldBindJSON(&editUserPreference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(editUserPreference)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.EditUserPreference(editUserPreference.UserID, editUserPreference.PreferenceID, editUserPreference.NewPreferenceName)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User preference edited successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) DeleteUserPreference(c *gin.Context) {
	var deleteUserPreference models.DeleteUserPreferenceRequest
	if err := c.ShouldBindJSON(&deleteUserPreference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(deleteUserPreference)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.DeleteUserPreference(deleteUserPreference.UserID, deleteUserPreference.PreferenceID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User preference deleted successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) GetUserInterests(c *gin.Context) {
	userID := c.Param("user_id")
	id, err := strconv.Atoi(userID)
	interests, err := ur.GRPC_Client.GetUserInterests(uint64(id))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to fetch user interests", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User interests fetched successfully", interests, nil)
	c.JSON(http.StatusOK, success)
}

func (ur *UserHandler) GetUserPreferences(c *gin.Context) {
	userID := c.Param("user_id")
	id, _ := strconv.Atoi(userID)
	preferences, err := ur.GRPC_Client.GetUserPreferences(uint64(id))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to fetch user preferences", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User preferences fetched successfully", preferences, nil)
	c.JSON(http.StatusOK, success)
}

// func (ur *UserHandler) SendMessage(c *gin.Context) {
// 	var requestData struct {
// 		Content    string         `json:"content"`
// 		Media      []models.Media `json:"media"`
// 		RecipentID int            `json:"user_id"`
// 	}
// 	if err := c.ShouldBindJSON(&requestData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
// 		return
// 	}

// 	// recipentID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
// 	// 	return
// 	// }
// 	authHeader := c.GetHeader("Authorization")
// 	token := helper.GetTokenFromHeader(authHeader)
// 	senderID, _, err := helper.ExtractUserIDFromToken(token)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 		return
// 	}

// 	// Prepare the message object
// 	message := models.UserMessage{
// 		SenderID:   uint(senderID),
// 		RecipentID: uint(requestData.RecipentID),
// 		Content:    requestData.Content,
// 		Media:      requestData.Media,
// 	}

// 	// Send the message
// 	_, err = ur.GRPC_Client.SendMessage(message)
// 	if err != nil {
// 		errorMessage := fmt.Sprintf("Failed to send message: %s", err.Error())
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
// 		return
// 	}

//		c.JSON(http.StatusCreated, message)
//	}

func (ur *UserHandler) SendMessageKafka(c *gin.Context) {
	fmt.Println("hii kafka")
	// Initialize Kafka producer
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Kafka producer"})
		return
	}
	defer producer.Close()

	var requestData struct {
		Content    string         `json:"content"`
		Media      []models.Media `json:"media"`
		RecipentID int            `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	// recipentID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
	// 	return
	// }
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)
	senderID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Prepare the message object
	message := models.UserMessage{
		SenderID:   uint(senderID),
		RecipentID: uint(requestData.RecipentID),
		Content:    requestData.Content,
		Media:      requestData.Media,
	}
	// Prepare message
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("failedddddddd")
		return
	}
	kafkaMessage := &sarama.ProducerMessage{
		Topic: "chat", // Adjust topic name here
		Value: sarama.StringEncoder(jsonData),
		// Add other message attributes as needed
	}

	// Send message to Kafka topic
	producer.Input() <- kafkaMessage

	c.JSON(http.StatusCreated, gin.H{"message": "Message sent to Kafka"})
}

func (ur *UserHandler) ReadMessages(c *gin.Context) {
	// Extract room ID from request context
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	messages, err := ur.GRPC_Client.ReadMessages(uint32(userID))
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to read messages: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, messages)
}

var (
	db            *gorm.DB
	upgrader      = websocket.Upgrader{}
	rooms         = make(map[string]*Room)
	connectionsMu sync.Mutex
)

type Message struct {
	SenderID   int64     `json:"sender_id"`
	ReceiverID int64     `json:"receiver_id"`
	Message    string    `json:"message"`
	Time       time.Time `json:"time"`
}

type User struct {
	UserID int64  `json:"user_id" gorm:"primary_key"`
	Name   string `json:"name"`
}

type Room struct {
	User1       int64
	User2       int64
	Connections []*websocket.Conn
	Ch          chan *Message
}

func (ur *UserHandler) handleWebSocket(c *gin.Context) {
	// Upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer ws.Close()

	// Get user ID from query parameters
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		return
	}

	// Get receiver ID from query parameters
	receiverID, err := strconv.ParseInt(c.Query("receiver_id"), 10, 64)
	if err != nil {
		log.Printf("Invalid receiver ID: %v", err)
		return
	}

	roomID := generateRoomID(userID, receiverID)

	// Create a new room if it doesn't exist
	connectionsMu.Lock()
	if _, ok := rooms[roomID]; !ok {
		rooms[roomID] = &Room{
			User1:       userID,
			User2:       receiverID,
			Connections: []*websocket.Conn{ws},
			Ch:          make(chan *Message),
		}
		go broadcastMessages(roomID)
	} else {
		rooms[roomID].Connections = append(rooms[roomID].Connections, ws)
	}
	connectionsMu.Unlock()

	// Fetch previous messages from the database
	prevMessages, err := getPreviousMessages(userID, receiverID)
	if err != nil {
		log.Printf("Error fetching previous messages: %v", err)
	} else {
		// Send previous messages to the client
		for _, msg := range prevMessages {
			if err := ws.WriteJSON(msg); err != nil {
				log.Printf("Error sending previous message to connection: %v", err)
			}
		}
	}

	// Remove connection when this function returns
	defer func() {
		connectionsMu.Lock()
		for i, conn := range rooms[roomID].Connections {
			if conn == ws {
				rooms[roomID].Connections = append(rooms[roomID].Connections[:i], rooms[roomID].Connections[i+1:]...)
				break
			}
		}
		connectionsMu.Unlock()
	}()

	// Listen for incoming messages
	for {
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		// Save message to database
		msg.Time = time.Now()
		userMessage := models.UserMessage{
			SenderID:   uint(msg.SenderID),
			RecipentID: uint(msg.ReceiverID),
			Content:    msg.Message,
			CreatedAt:  time.Now(),
		}
		if _, err := ur.GRPC_Client.SendMessage(userMessage); err != nil {
			log.Printf("Error saving message: %v", err)
			continue
		}

		// Send message to other user in the room
		rooms[roomID].Ch <- &msg
	}
}

func getPreviousMessages(userID, receiverID int64) ([]*Message, error) {
	var messages []*Message
	if err := db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", userID, receiverID, receiverID, userID).Order("time").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func generateRoomID(user1, user2 int64) string {
	// Sort user IDs to ensure consistency
	if user1 > user2 {
		user1, user2 = user2, user1
	}
	return strconv.FormatInt(user1, 10) + "-" + strconv.FormatInt(user2, 10)
}

func broadcastMessages(roomID string) {
	room := rooms[roomID]
	for msg := range room.Ch {
		for _, conn := range room.Connections {
			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("Error sending message to connection: %v", err)
			}
		}
	}
}
