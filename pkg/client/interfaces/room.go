package interfaces

import "github.com/bibin-zoz/api-gateway/pkg/utils/models"

type RoomClient interface {
	CreateRoom(roomData models.RoomData) (models.Room, error)
	EditRoom(roomData models.RoomData) (models.Room, error)
	ChangeRoomStatus(roomID uint32, status string) (models.Room, error)
	AddMembersToRoom(roomID uint32, userIds []uint32) (models.Room, error)
	GetRoomJoinRequests(roomID uint32) ([]models.RoomJoinRequest, error)
}
