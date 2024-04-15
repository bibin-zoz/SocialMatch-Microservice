package handlers

import (
	"net/http"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	GRPC_Client interfaces.UserClient
}

func NewUserHandler(UserClient interfaces.UserClient) *UserHandler {
	return &UserHandler{
		GRPC_Client: UserClient,
	}
}
func (ur *UserHandler) UserSignup(c *gin.Context) {
	var SignupDetail models.UserSignup
	if err := c.ShouldBindJSON(&SignupDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(SignupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UsersSignUp(SignupDetail)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}
func (ur *UserHandler) Userlogin(c *gin.Context) {
	// var UserLoginDetail models.UserLogin
	// if err := c.ShouldBindJSON(&UserLoginDetail); err != nil {
	// 	errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errs)
	// }
	// err := validator.New().Struct(UserLoginDetail)
	// if err != nil {
	// 	errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errs)
	// 	return
	// }
	// user, err := ur.GRPC_Client.UserLogin(UserLoginDetail)
	// if err != nil {
	// 	errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errs)
	// 	return
	// }
	success := response.ClientResponse(http.StatusCreated, "User successfully logged in with password", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (ur *UserHandler) UserEditDetails(c *gin.Context) {
	var EditDetails models.UserSignup
	if err := c.ShouldBindJSON(&EditDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(EditDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UserEditDetails(EditDetails)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User Details edited successfully", user, nil)
	c.JSON(http.StatusCreated, success)
}
