package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
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

func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
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

func (ad *AdminHandler) GetPreferences(c *gin.Context) {
	preferences, err := ad.GRPC_Client.GetPreferences()
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch preferences", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Preferences fetched successfully", preferences, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AdminHandler) AddInterest(c *gin.Context) {
	var interest models.AddInterest

	if err := c.ShouldBindJSON(&interest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
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

func (ad *AdminHandler) EditInterest(c *gin.Context) {
	var interest models.Intrests
	if err := c.ShouldBindJSON(&interest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	_, err := ad.GRPC_Client.EditInterest(interest.ID, interest.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to edit interest", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Interest edited successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

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

func (ad *AdminHandler) AddPreference(c *gin.Context) {
	var preference models.AddPreferences
	if err := c.ShouldBindJSON(&preference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
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

func (ad *AdminHandler) EditPreference(c *gin.Context) {
	var preference models.Preferences
	if err := c.ShouldBindJSON(&preference); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	_, err := ad.GRPC_Client.EditPreference(preference.ID, preference.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to edit preference", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Preference edited successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

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
