package handler

import (
	"net/http"
	"trinity/internal/service"
	"trinity/pkg/validator"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{authService}
}

func (h *authHandler) Login(c *gin.Context) {
	var request service.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request",
			"status": http.StatusBadRequest})
		return
	}
	if err := validator.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}
	accessToken, err := h.authService.Login(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(),
			"status": http.StatusUnauthorized})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login Successfully",
		"tokens": gin.H{"access_token": accessToken}})
}
