package handlers

import (
	"net/http"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	models "github.com/bibin-zoz/api-gateway/pkg/utils/models"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAuthHandler struct {
	GRPC_Client interfaces.UserClient
}

func NewUserAuthHandler(UserClient interfaces.UserClient) *UserAuthHandler {
	return &UserAuthHandler{
		GRPC_Client: UserClient,
	}
}

// UserSignup godoc
// @Summary User Signup
// @Description User signup with details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserSignup true "Signup details"
// @Success 201 {object} models.User
// @Failure 400 {object} response.Response
// @Router /users/signup [post]
func (ur *UserAuthHandler) UserSignup(c *gin.Context) {
	var SignupDetail models.UserSignup
	if err := c.ShouldBindJSON(&SignupDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to bind Details not in correct format", nil, nil)
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

// Userlogin godoc
// @Summary User Login
// @Description User login with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "Login details"
// @Success 201 {object} models.User
// @Failure 400 {object} response.Response
// @Router /users/login [post]
func (ur *UserAuthHandler) Userlogin(c *gin.Context) {
	var UserLoginDetail models.UserLogin
	if err := c.ShouldBindJSON(&UserLoginDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}
	err := validator.New().Struct(UserLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UserLogin(UserLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully logged in with password", user, nil)
	c.JSON(http.StatusCreated, success)
}

// UserOtpReq godoc
// @Summary Request OTP
// @Description Request OTP for user verification
// @Tags users
// @Accept json
// @Produce json
// @Param request body interface{} true "OTP request details"
// @Success 201 {object} models.User
// @Failure 400 {object} response.Response
// @Router /users/otp-request [post]
func (ur *UserAuthHandler) UserOtpReq(c *gin.Context) {
	var EmailOtp models.UserVerificationRequest
	if err := c.ShouldBindJSON(&EmailOtp); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Email not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(EmailOtp)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ur.GRPC_Client.UserOtpRequest(EmailOtp)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Otp sented successfully", user, nil)
	c.JSON(http.StatusCreated, success)
}

// UserOtpVerification godoc
// @Summary Verify OTP
// @Description Verify user OTP
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.Otp true "OTP verification details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/otp-verify [post]
func (ur *UserAuthHandler) UserOtpVerification(c *gin.Context) {
	var EmailOtp models.Otp
	if err := c.ShouldBindJSON(&EmailOtp); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Email not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(EmailOtp)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	_, err = ur.GRPC_Client.UserOtpVerificationReq(EmailOtp)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "otp verified successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}
