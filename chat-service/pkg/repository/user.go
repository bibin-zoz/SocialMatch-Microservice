package repository

import (
	"errors"
	"fmt"

	"github.com/bibin-zoz/social-match-chat-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-chat-svc/pkg/repository/interface"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/utils/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type chatRepository struct {
	DB          *gorm.DB
	MongoClient *mongo.Client
}

func NewConnectionRepository(DB *gorm.DB, mongoClient *mongo.Client) interfaces.ChatRepository {
	return &chatRepository{
		DB:          DB,
		MongoClient: mongoClient,
	}

}
func (ur *chatRepository) SaveMessage(message *domain.UserMessage) (uint, error) {
	result := ur.DB.Create(message)
	if result.Error != nil {
		return 0, result.Error
	}
	return message.ID, nil
}

func (ur *chatRepository) SaveMedia(media *domain.Media) error {
	fmt.Println("repohiii")
	fmt.Println(media)
	result := ur.DB.Create(media)

	if result.Error != nil {
		return result.Error
	}
	fmt.Println("nill")
	return nil
}
func (ur *chatRepository) GetMessages(receiverID, senderID uint64) ([]models.UserMessage, error) {
	var messages []models.UserMessage

	// Preload media for each message
	if err := ur.DB.Preload("Media").Where("(recipent_id = ? AND sender_id = ?) OR ( recipent_id = ? AND sender_id = ?)", receiverID, senderID, senderID, receiverID).Find(&messages).Error; err != nil {
		return nil, errors.New("error retrieving messages")
	}

	return messages, nil
}
