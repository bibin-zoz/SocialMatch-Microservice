package interfaces

import (
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
)

type UserUseCase interface {
	UsersSignUp(user models.UserSignup) (domain.TokenUser, error)
	UsersLogin(user models.UserLogin) (domain.TokenUser, error)
	UserEditDetails(user models.UserSignup) (models.UserDetails, error)
	UserGenerateOtp(email string) (string, error)
	UserVerifyOtp(otp string, email string) (bool, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUserStatus(id int) (models.UserDetail, error)
	AddUserInterest(userID uint64, interestID int) error
	AddUserPreference(userID uint64, preferenceID int) error
	EditUserInterest(userID uint64, interestID uint64, newInterestName string) error
	EditUserPreference(userID uint64, preferenceID uint64, newPreferenceName string) error
	DeleteUserInterest(userID uint64, interestID uint64) error
	DeleteUserPreference(userID uint64, preferenceID uint64) error
	GetUserInterests(userID uint64) ([]string, error)
	GetUserPreferences(userID uint64) ([]string, error)
	FollowUser(senderID, userID int64) error
	BlockConnection(senderID, userID int64) error
	SendMessage(message *models.UserMessage) error
	SendMessageKafka(message *models.UserMessage) error
}
