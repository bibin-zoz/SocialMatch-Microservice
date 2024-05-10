package interfaces

import (
	"github.com/bibin-zoz/social-match-room-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-room-svc/pkg/utils/models"
)

type RoomUseCase interface {
	CreateRoom(room domain.Room) (*domain.Room, error)
	EditRoom(roomEdit models.RoomEdit) (*domain.Room, error)
	ChangeStatus(roomID uint, status string) (*domain.Room, error)
	AddMembers(roomID uint, userIDs []uint32) (*domain.Room, error)
	GetAllRooms() ([]*domain.Room, error)
	SeeRoomJoinRequests(roomID uint32) ([]models.RoomJoinRequest, error)
	GetRoomMembers(roomID uint32) ([]models.RoomMember, error)
	GetMessages(roomID uint) ([]*domain.Message, error)
	SendMessage(message domain.Message) (*domain.Message, error)
}
