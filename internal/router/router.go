package router

import (
	"net/http"
	"trinity/config"
	"trinity/internal/handler"
	"trinity/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler handler.UserHandler,
	authHandler handler.AuthHandler,
	voucherHandler handler.VoucherHandler,
	campaignHandler handler.CampaignHandler,
	cfg *config.Config,
) *gin.Engine {
	router := gin.Default()

	// Initialize cors config
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))

	apiGroup := router.Group("/api", PreflightHandler())

	userGroup := apiGroup.Group("/user")
	{
		userGroup.POST("/create-user", userHandler.CreateUser)
		userGroup.GET("/get-user", middleware.AuthMiddleware(), userHandler.GetUser)
	}

	voucherGroup := apiGroup.Group("/voucher")
	{
		voucherGroup.POST("/create-voucher", middleware.AuthMiddleware(), voucherHandler.CreateVoucher)
		voucherGroup.GET("/get-voucher", middleware.AuthMiddleware(), voucherHandler.GetVoucherByVoucherName)
	}

	campaignGroup := apiGroup.Group("/campaign")
	{
		campaignGroup.POST("/create-campaign", middleware.AuthMiddleware(), campaignHandler.CreateCampaign)
		campaignGroup.GET("/get-campaign", middleware.AuthMiddleware(), campaignHandler.GetCampaignByLink)
	}

	authGroup := apiGroup.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}

	apiGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong v1.1")
	})

	return router
}

func PreflightHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
