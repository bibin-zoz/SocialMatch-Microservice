package interfaces

import (
	"github.com/bibin-zoz/social-match-connections-svc/pkg/utils/models"
)

type ConnectionRepository interface {
	AddConnection(senderID, userID uint) error
	CheckFriends(senderID, userID uint) (bool, error)
	AddConnectionRequest(senderID, userID uint) error
	BlockConnection(senderID, userID uint) error
	GetConnections(userID uint) ([]models.UserDetails, error)
	CheckConnectionRequestExist(senderID, userID uint) (bool, error)
	CheckUserConnection(userID int64, recvID int64) (bool, error)
}
