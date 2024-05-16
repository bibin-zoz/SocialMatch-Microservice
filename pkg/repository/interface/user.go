package interfaces

import (
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
)

type UserRepository interface {
	CheckUserExistsByEmail(email string) (*domain.User, error)
	CheckUserExistsByPhone(phone string) (*domain.User, error)
	UserSignUp(user models.UserSignup) (models.UserDetails, error)
	FindUserByEmail(user models.UserLogin) (models.UserDetail, error)
	UserEditDetails(user models.UserSignup) (models.UserDetails, error)
	FetchAllUsers() ([]domain.User, error)
	GetUserByID(id int) (domain.User, error)
	BlockUser(userID uint) error
	UnblockUser(userID uint) error
	CheckUserInterest(userID uint64, interestName string) (bool, error)
	CheckUserInterestByID(userID uint64, interestID uint64) (bool, error)
	AddUserInterest(userID uint64, interestID int) error
	EditUserInterest(userID uint64, interestID uint64, newInterestName string) error
	DeleteUserInterest(userID uint64, interestID uint64) error
	CheckUserPreference(userID uint64, preferenceName string) (bool, error)
	CheckUserPreferenceByID(userID uint64, preferenceID uint64) (bool, error)
	AddUserPreference(userID uint64, preferenceID int) error
	EditUserPreference(userID uint64, preferenceID uint64, newPreferenceName string) error
	DeleteUserPreference(userID uint64, preferenceID uint64) error
	GetUserPreferences(userID uint64) ([]string, error)
	GetUserInterests(userID uint64) ([]string, error)
	// CheckConnectionRequestExist(senderID, userID uint) (bool, error)
	AddConnection(senderID, userID uint) error
	CheckFriends(senderID, userID uint) (bool, error)
	// AddConnectionRequest(senderID, userID uint) error
	// BlockConnection(senderID, userID uint) error
	SaveMessage(message *domain.UserMessage) (uint, error)
	SaveMedia(media *domain.Media) error
	GetMessages(receiverID, senderID uint64) ([]models.UserMessage, error)
	// GetConnections(userID uint) ([]models.UserDetails, error)
}
