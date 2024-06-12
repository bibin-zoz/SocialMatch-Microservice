package interfaces

import (
	"github.com/bibin-zoz/social-match-room-svc/pkg/domain"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateRoom(room domain.Room) (*domain.Room, error)
	GetRoomByID(roomID uint) (*domain.Room, error)
	UpdateRoom(room domain.Room) (*domain.Room, error)
	DeleteRoom(roomID uint) error
	AddRoomMember(roomID, userID uint) error
	GetAllRooms() ([]*domain.Room, error)
	SeeRoomJoinRequests(roomID uint) ([]*domain.RoomJoinReq, error)
	GetRoomMembers(roomID uint) ([]*domain.RoomMember, error)
	GetMessages(roomID uint) ([]*domain.Message, error)
	SaveMessage(message domain.Message) (*domain.Message, *gorm.DB, error)
	SaveMedia(tx *gorm.DB, messageID uint, media []domain.Media) ([]*domain.Media, error)
	GetRoomMembersById(roomID uint, userID uint) (*domain.RoomMember, error)
	IsUserConnectedToRoom(userID, roomID uint) (bool, error)
	CloseExpiredRooms() error
}
