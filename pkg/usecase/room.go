package usecase

import (
	"errors"
	"fmt"

	"github.com/bibin-zoz/social-match-room-svc/pkg/config"
	"github.com/bibin-zoz/social-match-room-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-room-svc/pkg/repository/interface"
	services "github.com/bibin-zoz/social-match-room-svc/pkg/usecase/interface"
	"github.com/bibin-zoz/social-match-room-svc/pkg/utils/models"
)

type roomUseCase struct {
	roomRepository interfaces.RoomRepository
	Config         config.Config
}

func NewRoomUseCase(repository interfaces.RoomRepository, config config.Config) services.RoomUseCase {
	return &roomUseCase{
		roomRepository: repository,
		Config:         config,
	}
}

func (uc *roomUseCase) CreateRoom(room domain.Room) (*domain.Room, error) {
	fmt.Println("usecase", room)
	if room.RoomName == "" || len(room.RoomName) < 6 {
		return nil, errors.New("room name is required and atleast 6 characters is required")
	}

	if room.MaxPersons <= 0 {
		return nil, errors.New("max persons must be greater than 0")
	}

	createdRoom, err := uc.roomRepository.CreateRoom(room)
	if err != nil {
		return nil, err
	}

	return createdRoom, nil
}

func (uc *roomUseCase) EditRoom(roomEdit models.RoomEdit) (*domain.Room, error) {

	existingRoom, err := uc.roomRepository.GetRoomByID(roomEdit.RoomID)
	if err != nil {
		return nil, err
	}

	existingRoom.RoomName = roomEdit.RoomName
	existingRoom.Description = roomEdit.Description
	existingRoom.MaxPersons = roomEdit.MaxPersons

	updatedRoom, err := uc.roomRepository.UpdateRoom(*existingRoom)
	if err != nil {
		return nil, err
	}

	return updatedRoom, nil
}

func (uc *roomUseCase) ChangeStatus(roomID uint, status string) (*domain.Room, error) {
	if status != "public" && status != "private" {
		return nil, errors.New("invalid status")
	}

	existingRoom, err := uc.roomRepository.GetRoomByID(roomID)
	if err != nil {
		return nil, err
	}

	existingRoom.Status = status

	updatedRoom, err := uc.roomRepository.UpdateRoom(*existingRoom)
	if err != nil {
		return nil, err
	}

	return updatedRoom, nil
}
func (uc *roomUseCase) AddMembers(roomID uint, userIDs []uint32) (*domain.Room, error) {
	room, err := uc.roomRepository.GetRoomByID(roomID)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, errors.New("room not found")
	}

	for _, userID := range userIDs {
		// user, err := uc.userRepository.GetUserByID(userID)
		// if err != nil {
		// 	return nil, err
		// }
		// if user == nil {
		// 	return nil, fmt.Errorf("user with ID %d not found", userID)
		// }

		err = uc.roomRepository.AddRoomMember(roomID, uint(userID))
		if err != nil {
			return nil, err
		}
	}

	return room, nil
}
func (uc *roomUseCase) GetAllRooms() ([]*domain.Room, error) {
	rooms, err := uc.roomRepository.GetAllRooms()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
func (uc *roomUseCase) SeeRoomJoinRequests(roomID uint32) ([]models.RoomJoinRequest, error) {
	joinRequests, err := uc.roomRepository.SeeRoomJoinRequests(uint(roomID))
	if err != nil {
		return nil, err
	}

	var roomJoinRequests []models.RoomJoinRequest
	for _, req := range joinRequests {
		roomJoinRequests = append(roomJoinRequests, models.RoomJoinRequest{
			UserID: uint32(req.UserID),
			// Username: req.Username,
		})
	}

	return roomJoinRequests, nil
}

func (uc *roomUseCase) GetRoomMembers(roomID uint32) ([]models.RoomMember, error) {
	members, err := uc.roomRepository.GetRoomMembers(uint(roomID))
	if err != nil {
		return nil, err
	}

	var roomMembers []models.RoomMember
	for _, member := range members {
		roomMembers = append(roomMembers, models.RoomMember{
			UserID: member.UserID,
			// Username: member.Username,
		})
	}

	return roomMembers, nil
}
func (uc *roomUseCase) SendMessage(message domain.Message) (*domain.Message, error) {

	// Validate message content
	if message.Content == "" {
		return nil, errors.New("message content cannot be empty")
	}

	// Retrieve room members
	roomMembers, err := uc.roomRepository.GetRoomMembersById(uint(message.RoomID), uint(message.UserID))
	if err != nil || roomMembers == nil {
		return nil, errors.New("user is not a member of the room")
	}

	// Call repository method to send message
	sentMessage, tx, err := uc.roomRepository.SaveMessage(message)
	if err != nil {
		return nil, err
	}

	if message.Media != nil {
		saveMedia, err := uc.roomRepository.SaveMedia(tx, sentMessage.ID, message.Media)
		if err != nil || saveMedia == nil {
			tx.Rollback()
			return nil, err
		}

	}

	// Rollback transaction if error occurs
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return sentMessage, nil
}

func (uc *roomUseCase) GetMessages(roomID uint) ([]*domain.Message, error) {
	// Call repository method to get messages for the specified room
	messages, err := uc.roomRepository.GetMessages(roomID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
