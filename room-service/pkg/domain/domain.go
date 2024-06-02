package domain

import (
	"time"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
)

type User struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
	Phone     string `json:"phone" gorm:"validate:required"`
	Blocked   bool   `json:"blocked" gorm:"default:false;validate:required"`
	Username  string `json:"user_name"`
	GenderID  int    `json:"gender_id"`
	Age       int    `json:"age"`
}

type TokenUser struct {
	User         models.UserDetails
	AccessToken  string
	RefreshToken string
}

// Room represents a room entity
type Room struct {
	ID           uint       `gorm:"primaryKey"`
	RoomName     string     `gorm:"not null"`
	CreatedDate  time.Time  `gorm:"not null"`
	CloseDate    *time.Time `gorm:"default:null"`
	Description  string
	MaxPersons   uint          `gorm:"not null"`
	CreatorID    int           `gorm:"not null"`
	Status       string        `gorm:"check:status IN ('public', 'private')"`
	Preferences  []Preference  `gorm:"many2many:room_preferences;"`
	RoomMembers  []RoomMember  `gorm:"foreignKey:RoomID"`
	RoomJoinReqs []RoomJoinReq `gorm:"foreignKey:RoomID"`
}

type Preference struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type RoomMember struct {
	RoomID uint
	UserID uint
}

type RoomJoinReq struct {
	RoomID uint
	UserID uint
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
	ID        uint   `json:"id"`
	MessageID uint   `gorm:"foreignKey:MessageID"`
	Filename  string `json:"filename"`
}
