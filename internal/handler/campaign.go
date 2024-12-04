package handler

import (
	"net/http"
	"trinity/internal/service"
	"trinity/pkg/validator"

	"github.com/gin-gonic/gin"
)

type CampaignHandler interface {
	CreateCampaign(c *gin.Context)
	GetCampaignByLink(c *gin.Context)
}

type campaignHandler struct {
	campaignService service.CampaignService
}

func NewCampaignHandler(campaignService service.CampaignService) CampaignHandler {
	return &campaignHandler{campaignService}
}

func (handler *campaignHandler) CreateCampaign(c *gin.Context) {
	var request service.CreateCampaignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body",
			"status": http.StatusBadRequest})
		return
	}

	// email := request.Email
	// emailAuthen := c.MustGet("email")

	// if emailAuthen != email {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token",
	// 		"status": http.StatusUnauthorized})
	// 	return
	// }

	if err := validator.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	data, err := handler.campaignService.CreateCampaign(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (handler *campaignHandler) GetCampaignByLink(c *gin.Context) {
	var request service.GetCampaignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body",
			"status": http.StatusBadRequest})
		return
	}

	if err := validator.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	data, err := handler.campaignService.GetCampaignByLink(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get Campaign Successfully",
		"data": data})
}
