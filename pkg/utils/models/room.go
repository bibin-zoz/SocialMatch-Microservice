package models

import "time"

// Room represents a room entity
type Room struct {
	ID          uint32   `json:"id"`
	RoomName    string   `json:"room_name"`
	Description string   `json:"description"`
	MaxPersons  uint32   `json:"max_persons"`
	Status      string   `json:"status"`
	Preferences []string `json:"preferences"`
}

// RoomJoinRequest represents a request to join a room
type RoomJoinRequest struct {
	UserID   uint32 `json:"user_id"`
	Username string `json:"username"`
}

// RoomData represents data for creating or editing a room
type RoomData struct {
	ID          uint32   `json:"id,omitempty"`
	RoomName    string   `json:"room_name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	MaxPersons  uint32   `json:"max_persons" validate:"required,min=1"`
	Status      string   `json:"status" validate:"required,oneof=public private"`
	Preferences []string `json:"preferences"`
}

type RoomMember struct {
	UserID   uint32 `json:"user_id"`
	Username string `json:"username"`
}
type Message struct {
	ID        uint `gorm:"primaryKey"`
	RoomID    uint
	UserID    uint
	Content   string
	CreatedAt time.Time
	Media     []Media `json:"media"`
}
type Media struct {
	ID       uint   `json:"id"`
	Filename string `json:"filename"`
}
