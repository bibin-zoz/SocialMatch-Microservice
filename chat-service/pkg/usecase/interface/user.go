package interfaces

import (
	"github.com/bibin-zoz/social-match-chat-svc/pkg/utils/models"
)

type ChatUseCase interface {
	SendMessage(message *models.UserMessage) error
	ConsumeAndProcessMessages() 
	GetMessages(receiverID, senderID uint64) ([]models.UserMessage, error) 
}
