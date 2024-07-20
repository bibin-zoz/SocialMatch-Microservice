package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/helper"
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

// UserEditDetails godoc
// @Summary Edit User Details
// @Description Edit user details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserUpdateDetails true "Edit details"
// @Success 201 {object} models.User
// @Failure 400 {object} response.Response
// @Router /users/edit [put]
func (ur *UserHandler) UserEditDetails(c *gin.Context) {
	var EditDetails models.UserUpdateDetails
	fmt.Println("edit", EditDetails)
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

// UpdateProfilePhoto godoc
// @Summary Update Profile Photo
// @Description Update user's profile photo
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param image_paths formData []file true "Profile photos"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/profile-photo [post]
func (ur *UserHandler) UpdateProfilePhoto(c *gin.Context) {
	type UpdateProfilePhotoForm struct {
		ImagePaths []string `form:"image_paths"`
	}
	form, err := c.MultipartForm()
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid form data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	files := form.File["images"]
	fmt.Println("doneeeeeeeeeeeee", files)
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)

	userID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 4. Call gRPC client to update profile photo
	err = ur.GRPC_Client.UpdateProfilePhoto(int64(userID), files)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update profile photo", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	// 5. Handle success and return response
	success := response.ClientResponse(http.StatusOK, "Profile photo updated successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// GetAllUsers godoc
// @Summary Get All Users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Failure 400 {object} response.Response
// @Router /users [get]
func (ur *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := ur.GRPC_Client.GetAllUsers()
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Users fetched successfully", users, nil)
	c.JSON(http.StatusCreated, success)

}

// AddUserInterest godoc
// @Summary Add User Interest
// @Description Add user interest
// @Tags users
// @Accept json
// @Produce json
// @Param user_id body models.AddUserInterestRequest true "Add interest request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/add-interest [post]
func (ur *UserHandler) AddUserInterest(c *gin.Context) {
	var addUserInterest models.AddUserInterestRequest
	if err := c.ShouldBindJSON(&addUserInterest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(addUserInterest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.AddUserInterest(addUserInterest.UserID, addUserInterest.InterestID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User interest added successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

// EditUserInterest godoc
// @Summary Edit User Interest
// @Description Edit user interest
// @Tags users
// @Accept json
// @Produce json
// @Param user_id body models.EditUserInterestRequest true "Edit interest request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/edit-interest [put]
func (ur *UserHandler) EditUserInterest(c *gin.Context) {
	var editUserInterest models.EditUserInterestRequest
	if err := c.ShouldBindJSON(&editUserInterest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(editUserInterest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.EditUserInterest(editUserInterest.UserID, editUserInterest.InterestID, editUserInterest.NewInterestName)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User interest edited successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

// DeleteUserInterest godoc
// @Summary Delete User Interest
// @Description Delete user interest
// @Tags users
// @Accept json
// @Produce json
// @Param user_id body models.DeleteUserInterestRequest true "Delete interest request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/delete-interest [delete]
func (ur *UserHandler) DeleteUserInterest(c *gin.Context) {
	var deleteUserInterest models.DeleteUserInterestRequest
	if err := c.ShouldBindJSON(&deleteUserInterest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(deleteUserInterest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.DeleteUserInterest(deleteUserInterest.UserID, deleteUserInterest.InterestID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User interest deleted successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

// AddUserPreference godoc
// @Summary Add User Preference
// @Description Add user preference
// @Tags users
// @Accept json
// @Produce json
// @Param user_id body models.AddUserPreferenceRequest true "Add preference request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/add-preference [post]
func (ur *UserHandler) AddUserPreference(c *gin.Context) {
	var addUserPreference models.AddUserPreferenceRequest
	if err := c.ShouldBindJSON(&addUserPreference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(addUserPreference)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.AddUserPreference(addUserPreference.UserID, addUserPreference.PreferenceID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User preference added successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

// EditUserPreference godoc
// @Summary Edit User Preference
// @Description Edit user preference
// @Tags users
// @Accept json
// @Produce json
// @Param user_id body models.EditUserPreferenceRequest true "Edit preference request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/edit-preference [put]
func (ur *UserHandler) EditUserPreference(c *gin.Context) {
	var editUserPreference models.EditUserPreferenceRequest
	if err := c.ShouldBindJSON(&editUserPreference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(editUserPreference)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.EditUserPreference(editUserPreference.UserID, editUserPreference.PreferenceID, editUserPreference.NewPreferenceName)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User preference edited successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

// DeleteUserPreference godoc
// @Summary Delete User Preference
// @Description Delete user preference
// @Tags users
// @Accept json
// @Produce json
// @Param user_id body models.DeleteUserPreferenceRequest true "Delete preference request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/delete-preference [delete]
func (ur *UserHandler) DeleteUserPreference(c *gin.Context) {
	var deleteUserPreference models.DeleteUserPreferenceRequest
	if err := c.ShouldBindJSON(&deleteUserPreference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(deleteUserPreference)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.DeleteUserPreference(deleteUserPreference.UserID, deleteUserPreference.PreferenceID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User preference deleted successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

// @Summary Get user interests
// @Description Get a list of user interests
// @Tags users
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} []string "User interests"
// @Failure 400 {object} response.Response
// @Router /users/{user_id}/interests [get]
func (ur *UserHandler) GetUserInterests(c *gin.Context) {
	userID := c.Param("user_id")
	id, err := strconv.Atoi(userID)
	interests, err := ur.GRPC_Client.GetUserInterests(uint64(id))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to fetch user interests", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User interests fetched successfully", interests, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Get user preferences
// @Description Get a list of user preferences
// @Tags users
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} []models.Preferences "User preferences"
// @Failure 400 {object} response.Response
// @Router /users/{user_id}/preferences [get]
func (ur *UserHandler) GetUserPreferences(c *gin.Context) {
	userID := c.Param("user_id")
	id, _ := strconv.Atoi(userID)
	preferences, err := ur.GRPC_Client.GetUserPreferences(uint64(id))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to fetch user preferences", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User preferences fetched successfully", preferences, nil)
	c.JSON(http.StatusOK, success)
}

func (ur *UserHandler) ReadMessages(c *gin.Context) {
	// Extract room ID from request context
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)

	// Extract user ID from token
	userID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var req struct {
		UserID int `json:"user_id"`
	}
	// Bind the JSON request body to the req struct
	if err := c.ShouldBindJSON(&req); err != nil {
		// If binding fails, return a 400 Bad Request response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	messages, err := ur.GRPC_Client.ReadMessages(uint32(userID), uint32(req.UserID))
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to read messages: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", messages, nil)

	c.JSON(http.StatusOK, success)
}

// func (ur *UserHandler) AddProfilePhoto(c *gin.Context) {
// 	type UpdateProfilePhotoForm struct {
// 		ImagePaths []string `form:"image_paths"`
// 	}
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		errs := response.ClientResponse(http.StatusBadRequest, "Invalid form data", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errs)
// 		return
// 	}

// 	files := form.File["images"]
// 	fmt.Println("doneeeeeeeeeeeee", files)
// 	authHeader := c.GetHeader("Authorization")
// 	token := helper.GetTokenFromHeader(authHeader)

// 	userID, _, err := helper.ExtractUserIDFromToken(token)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 		return
// 	}

// 	// 4. Call gRPC client to update profile photo
// 	err = ur.GRPC_Client.AddProfilePhoto(int64(userID), files)
// 	if err != nil {
// 		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update profile photo", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errs)
// 		return
// 	}

// 	// 5. Handle success and return response
// 	success := response.ClientResponse(http.StatusOK, "Profile photo updated successfully", nil, nil)
// 	c.JSON(http.StatusOK, success)
// }
