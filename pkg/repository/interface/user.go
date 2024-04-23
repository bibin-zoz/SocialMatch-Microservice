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
	AddUserInterest(userID uint64, interestName string) error
	EditUserInterest(userID uint64, interestID uint64, newInterestName string) error
	DeleteUserInterest(userID uint64, interestID uint64) error
}
