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
	AddUserInterest(userID uint64, interestName string) error
	EditUserInterest(userID uint64, interestID uint64, newInterestName string) error
	DeleteUserInterest(userID uint64, interestID uint64) error
}
