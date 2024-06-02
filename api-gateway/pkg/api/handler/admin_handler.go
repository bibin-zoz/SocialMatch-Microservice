package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminHandler struct {
	GRPC_Client      interfaces.AdminClient
	GRPC_USER_Client interfaces.UserClient
}

func NewAdminHandler(adminClient interfaces.AdminClient, userClient interfaces.UserClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client:      adminClient,
		GRPC_USER_Client: userClient,
	}

}

// LoginHandler godoc
// @Summary Admin Login
// @Description Authenticate admin user
// @Tags admin
// @Accept json
// @Produce json
// @Param adminDetails body models.AdminLogin true "Admin Login Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/login [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid details", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	fmt.Println("admin", admin)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	fmt.Println("success")
	success := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}

// UpdateUserStatus godoc
// @Summary Update User Status
// @Description Update the status of a user
// @Tags admin
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/users [put]
func (ad *AdminHandler) UpdateUserStatus(c *gin.Context) {
	userID := c.Query("id")
	fmt.Println("userid", userID)
	id, _ := strconv.Atoi(userID)
	_, err := ad.GRPC_USER_Client.UpdateUserStatus(int64(id))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to update user status", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Updated user successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}
func (ad *AdminHandler) GetInterests(c *gin.Context) {
	interests, err := ad.GRPC_Client.GetIntrests()
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch interests", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Interests fetched successfully", interests, nil)
	c.JSON(http.StatusOK, success)
}

// GetInterests godoc
// @Summary Get Interests
// @Description Get all interests
// @Tags admin
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/interests [get]
func (ad *AdminHandler) GetPreferences(c *gin.Context) {
	fmt.Println("hiii")
	preferences, err := ad.GRPC_Client.GetPreferences()
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch preferences", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Preferences fetched successfully", preferences, nil)
	c.JSON(http.StatusOK, success)
}

// GetPreferences godoc
// @Summary Get Preferences
// @Description Get all preferences
// @Tags admin
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/preferences [get]
func (ad *AdminHandler) AddInterest(c *gin.Context) {
	var interest models.AddInterest

	if err := c.ShouldBindJSON(&interest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if err := validator.New().Struct(interest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid details", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	fmt.Println("interesdddddddddddddddt", interest)
	_, err := ad.GRPC_Client.AddInterest(interest.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to add interest", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Interest added successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// AddInterest godoc
// @Summary Add Interest
// @Description Add a new interest
// @Tags admin
// @Accept json
// @Produce json
// @Param interest body models.AddInterest true "Interest to add"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/interests [post]
func (ad *AdminHandler) EditInterest(c *gin.Context) {
	var interest models.Intrests
	if err := c.ShouldBindJSON(&interest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if err := validator.New().Struct(interest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid details", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	_, err := ad.GRPC_Client.EditInterest(int(interest.ID), interest.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to edit interest", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Interest edited successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// EditInterest godoc
// @Summary Edit Interest
// @Description Edit an existing interest
// @Tags admin
// @Accept json
// @Produce json
// @Param interest body models.Intrests true "Interest to edit"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/interests [put]
func (ad *AdminHandler) DeleteInterest(c *gin.Context) {
	id := c.Param("id")
	interestID, err := strconv.Atoi(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid interest ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = ad.GRPC_Client.DeleteInterest((interestID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete interest", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Interest deleted successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// DeleteInterest godoc
// @Summary Delete Interest
// @Description Delete an existing interest
// @Tags admin
// @Produce json
// @Param id path int true "Interest ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/interests/{id} [delete]
func (ad *AdminHandler) AddPreference(c *gin.Context) {
	var preference models.AddPreferences
	if err := c.ShouldBindJSON(&preference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if err := validator.New().Struct(preference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid details", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	_, err := ad.GRPC_Client.AddPreference(preference.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to add preference", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Preference added successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// AddPreference godoc
// @Summary Add Preference
// @Description Add a new preference
// @Tags admin
// @Accept json
// @Produce json
// @Param preference body models.AddPreferences true "Preference to add"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/preferences [post]
func (ad *AdminHandler) EditPreference(c *gin.Context) {
	var preference models.Preferences
	if err := c.ShouldBindJSON(&preference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if err := validator.New().Struct(preference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid details", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	_, err := ad.GRPC_Client.EditPreference(int(preference.ID), preference.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to edit preference", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Preference edited successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// EditPreference godoc
// @Summary Edit Preference
// @Description Edit an existing preference
// @Tags admin
// @Accept json
// @Produce json
// @Param preference body models.Preferences true "Preference to edit"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/preferences [put]
func (ad *AdminHandler) DeletePreference(c *gin.Context) {
	id := c.Param("id")
	preferenceID, err := strconv.Atoi(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid preference ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = ad.GRPC_Client.DeletePreference((preferenceID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete preference", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Preference deleted successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}
