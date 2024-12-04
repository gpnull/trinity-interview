package handler

import (
	"net/http"
	"trinity/internal/service"
	"trinity/pkg/validator"

	"github.com/gin-gonic/gin"
)

type VoucherHandler interface {
	CreateVoucher(c *gin.Context)
	GetVoucherByVoucherName(c *gin.Context)
}

type voucherHandler struct {
	voucherService service.VoucherService
}

func NewVoucherHandler(voucherService service.VoucherService) VoucherHandler {
	return &voucherHandler{voucherService}
}

func (handler *voucherHandler) CreateVoucher(c *gin.Context) {
	var request service.CreateVoucherRequest
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

	data, err := handler.voucherService.CreateVoucher(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (handler *voucherHandler) GetVoucherByVoucherName(c *gin.Context) {
	var request service.GetVoucherRequest
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

	data, err := handler.voucherService.GetVoucherByVoucherName(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"status": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get Voucher Successfully",
		"data": data})
}
