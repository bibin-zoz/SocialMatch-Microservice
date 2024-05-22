package domain

import (
	"time"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
	"gorm.io/gorm"
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
type UserInterest struct {
	UserID     int
	InterestID int
}
type UserPreference struct {
	UserID       int
	PreferenceID int
}
type Interest struct {
	ID           int
	InterestName string
}

type Preference struct {
	ID             int
	PreferenceName string
}

type Connections struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	FriendID uint
	Status   string `gorm:"check:status IN ('pending', 'friends', 'blocked')"`
}

type Media struct {
	gorm.Model
	Message_id int    `gorm:foriegnKey"`
	Filename   string `json:"filename"`
}
type UserMessage struct {
	ID         uint `gorm:"primaryKey"`
	SenderID   uint
	RecipentID uint
	Content    string
	CreatedAt  time.Time
	Read       bool `dafault:"false"`
	// Media      []Media `json:"media"`
}
type UserProfilePhoto struct {
	UserID   int      `bson:"userId"`
	ImageURL []string `bson:"imageUrl"`
}
