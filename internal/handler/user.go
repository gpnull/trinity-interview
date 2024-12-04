package handler

import (
	"net/http"
	"trinity/internal/service"
	"trinity/pkg/validator"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
}

func (handler *userHandler) CreateUser(c *gin.Context) {
	var request service.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body",
			"status": http.StatusBadRequest})
		return
	}

	if c.Param("campaign_name") != "" {
		request.CampaignName = c.Param("campaign_name")
	}
	if err := validator.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	data, err := handler.userService.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (handler *userHandler) GetUser(c *gin.Context) {
	var request service.GetUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body",
			"status": http.StatusBadRequest})
		return
	}

	email := request.Email
	emailAuthen := c.MustGet("email")

	if emailAuthen != email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token",
			"status": http.StatusUnauthorized})
		return
	}

	if err := validator.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	data, err := handler.userService.GetUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get User Successfully",
		"data": data})
}
