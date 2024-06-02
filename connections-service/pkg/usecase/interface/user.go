package interfaces

import (
	"github.com/bibin-zoz/social-match-connections-svc/pkg/utils/models"
)

type ConnectionUseCase interface {
	FollowUser(senderID, userID int64) error
	BlockConnection(senderID, userID int64) error
	GetConnections(userID uint64) ([]models.UserDetails, error)
	CheckUserConnection(userID int64, recvID int64) (bool, error)
}
