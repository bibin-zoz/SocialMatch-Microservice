package interfaces

import "github.com/bibin-zoz/api-gateway/pkg/utils/models"

type UserClient interface {
	UsersSignUp(user models.UserSignup) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
	UserEditDetails(user models.UserSignup) (models.UserDetails, error)
	UserOtpRequest(user models.UserVerificationRequest) (models.Otp, error)
	UserOtpVerificationReq(user models.Otp) (models.UserDetail, error)
}
