package models

import (
	"time"
)

// Room represents a room in the system
type Room struct {
	ID           uint       `gorm:"primaryKey"`
	RoomName     string     `gorm:"not null"`
	CreatedDate  time.Time  `gorm:"not null"`
	CloseDate    *time.Time `gorm:"default:null"`
	Description  string
	MaxPersons   uint          `gorm:"not null"`
	Status       string        `gorm:"default:'public'"`
	Preferences  []Preference  `gorm:"many2many:room_preferences;"`
	RoomMembers  []RoomMember  `gorm:"foreignKey:RoomID"`
	RoomJoinReqs []RoomJoinReq `gorm:"foreignKey:RoomID"`
}
type UserDetails struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}

// Preference represents a preference in the system

type Preference struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

// RoomMember represents a member of a room
type RoomMember struct {
	RoomID   uint
	UserID   uint
	Username string `json:"username"`
}

// RoomJoinReq represents a request to join a room
type RoomJoinReq struct {
	RoomID uint
	UserID uint
}
type RoomEdit struct {
	RoomID      uint
	RoomName    string
	Description string
	MaxPersons  uint
	Status      string
}
type RoomJoinRequest struct {
	UserID   uint32 `json:"user_id"`
	Username string `json:"username"`
}
