package repository

import (
	"errors"
	"fmt"

	"github.com/bibin-zoz/social-match-room-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-room-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type roomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(DB *gorm.DB) interfaces.RoomRepository {
	return &roomRepository{
		DB: DB,
	}

}

func (rr *roomRepository) CreateRoom(room domain.Room) (*domain.Room, error) {
	err := rr.DB.Create(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (rr *roomRepository) GetRoomByID(roomID uint) (*domain.Room, error) {
	var room domain.Room
	err := rr.DB.First(&room, roomID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}

func (rr *roomRepository) UpdateRoom(room domain.Room) (*domain.Room, error) {
	err := rr.DB.Save(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (rr *roomRepository) DeleteRoom(roomID uint) error {
	err := rr.DB.Delete(&domain.Room{}, roomID).Error
	return err
}
func (rr *roomRepository) AddRoomMember(roomID, userID uint) error {
	member := domain.RoomMember{
		RoomID: roomID,
		UserID: userID,
	}
	if err := rr.DB.Create(&member).Error; err != nil {
		return err
	}

	return nil
}
func (rr *roomRepository) GetAllRooms() ([]*domain.Room, error) {
	var rooms []*domain.Room
	err := rr.DB.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
func (rr *roomRepository) GetRoomMembers(roomID uint) ([]*domain.RoomMember, error) {
	var members []*domain.RoomMember
	err := rr.DB.Where("room_id = ?", roomID).Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}
func (rr *roomRepository) GetRoomMembersById(roomID uint, userID uint) (*domain.RoomMember, error) {
	var user *domain.RoomMember
	err := rr.DB.Where("room_id = ? and user_id=?", userID, roomID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (rr *roomRepository) SeeRoomJoinRequests(roomID uint) ([]*domain.RoomJoinReq, error) {
	var requests []*domain.RoomJoinReq
	err := rr.DB.Where("room_id = ?", roomID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (r *roomRepository) SaveMessage(message domain.Message) (*domain.Message, *gorm.DB, error) {
	// Validate if the room exists
	room, err := r.GetRoomByID(message.RoomID)
	if err != nil {
		return nil, nil, err
	}
	if room == nil {
		return nil, nil, errors.New("room not found")
	}

	// Create the message in the database
	tx := r.DB.Begin()
	if err := tx.Create(&message).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	return &message, tx, nil
}

func (r *roomRepository) SaveMedia(tx *gorm.DB, messageID uint, media []domain.Media) ([]*domain.Media, error) {
	fmt.Println("SaveMedia")
	var savedMedia []*domain.Media

	// Save media in the database
	for _, m := range media {
		// Set the MessageID for each media
		m.MessageID = messageID

		if err := tx.Create(&m).Error; err != nil {
			return nil, err
		}
		savedMedia = append(savedMedia, &m)
	}

	return savedMedia, nil
}
func (r *roomRepository) GetMessages(roomID uint) ([]*domain.Message, error) {
	// Retrieve messages for the specified room
	var messages []*domain.Message
	err := r.DB.Where("room_id = ?", roomID).Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
