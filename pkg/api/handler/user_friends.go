package handlers

import (
	"net/http"

	"github.com/bibin-zoz/api-gateway/pkg/helper"
	response "github.com/bibin-zoz/api-gateway/pkg/utils/responce"
	"github.com/gin-gonic/gin"
)

func (ur *UserHandler) UserFollow(c *gin.Context) {
	var req struct {
		Userid int `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to bind Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)
	senderID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error fetching user details,make sure loginned", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.FollowUser(req.Userid, senderID)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "successfully followed user", nil, nil)
	c.JSON(http.StatusCreated, success)
}
func (ur *UserHandler) BlockUser(c *gin.Context) {
	var req struct {
		Userid int `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to bind Details not in correct format", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)
	senderID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error fetching user details,make sure loginned", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = ur.GRPC_Client.BlockUser(req.Userid, senderID)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to conenct to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "successfully Blocked user", req.Userid, nil)
	c.JSON(http.StatusCreated, success)
}
func (ur *UserHandler) GetConnections(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := helper.GetTokenFromHeader(authHeader)
	userID, _, err := helper.ExtractUserIDFromToken(token)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error fetching user details, make sure logged in", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	connections, err := ur.GRPC_Client.GetConnections(uint64(userID))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "failed to connect to server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "conenctions fetched successfully", connections, nil)

	c.JSON(http.StatusOK, success)
}
