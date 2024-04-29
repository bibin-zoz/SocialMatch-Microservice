package interfaces

import "github.com/bibin-zoz/api-gateway/pkg/utils/models"

type UserClient interface {
	UsersSignUp(user models.UserSignup) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
	UserEditDetails(user models.UserUpdateDetails) (models.UserDetails, error)
	UserOtpRequest(user models.UserVerificationRequest) (models.Otp, error)
	UserOtpVerificationReq(user models.Otp) (models.UserDetail, error)
	GetAllUsers() ([]models.Users, error)
	UpdateUserStatus(userid int64) (models.UserDetail, error)
	AddUserInterest(userID uint64, interestID int) error
	EditUserInterest(userID uint64, interestID uint64, newInterestName string) error
	DeleteUserInterest(userID uint64, interestID uint64) error
	AddUserPreference(userID uint64, preference_id int) error
	EditUserPreference(userID uint64, preferenceID uint64, newPreferenceName string) error
	DeleteUserPreference(userID uint64, preferenceID uint64) error
	GetUserInterests(userID uint64) ([]string, error)
	GetUserPreferences(userID uint64) ([]string, error)
}
