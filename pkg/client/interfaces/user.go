package interfaces

import "github.com/bibin-zoz/api-gateway/pkg/utils/models"

type UserClient interface {
	UsersSignUp(user models.UserSignup) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
}
