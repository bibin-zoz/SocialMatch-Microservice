package interfaces

import (
	"github.com/bibin-zoz/social-match-chat-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/utils/models"
)

type ChatRepository interface {
	SaveMessage(message *domain.UserMessage) (uint, error)
	SaveMedia(media *domain.Media) error
	GetMessages(receiverID, senderID uint64) ([]models.UserMessage, error)
}
