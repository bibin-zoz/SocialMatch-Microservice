package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/IBM/sarama"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	"github.com/gin-gonic/gin"
	// Update the import path as per your project structure
)

func GenerateRoomID(user1, user2 int64) string {
	// Sort user IDs to ensure consistency
	if user1 > user2 {
		user1, user2 = user2, user1
	}
	return strconv.FormatInt(user1, 10) + "-" + strconv.FormatInt(user2, 10)
}

func BroadcastMessages(roomID string, rooms map[string]*models.WebrtcRoom, connectionsMu *sync.Mutex) {
	room := rooms[roomID]
	for msg := range room.Ch {
		connectionsMu.Lock()
		for _, conn := range room.Connections {
			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("Error sending message to connection: %v", err)
			}
		}
		connectionsMu.Unlock()
	}
}
func SendMessageKafka(message models.UserMessage, c *gin.Context) error {
	fmt.Println("hii kafka")
	// Initialize Kafka producer
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Kafka producer"})
		return err
	}
	defer producer.Close()

	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("failedddddddd")
		return err
	}
	kafkaMessage := &sarama.ProducerMessage{
		Topic: "chat", // Adjust topic name here
		Value: sarama.StringEncoder(jsonData),
		// Add other message attributes as needed
	}

	// Send message to Kafka topic
	producer.Input() <- kafkaMessage

	// c.JSON(http.StatusCreated, gin.H{"message": "Message sent to Kafka"})
	return nil
}
