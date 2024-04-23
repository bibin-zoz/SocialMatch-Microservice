package domain

import "github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"

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
	UserID       int
	InterestName string
}
type UserPreference struct {
	UserID         int
	PreferenceName string
}
type Interest struct {
	ID           int
	InterestName string
}

type Preference struct {
	ID             int
	PreferenceName string
}
