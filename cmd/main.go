package main

import (
	"log"
	"trinity/config"
	"trinity/internal/handler"
	"trinity/internal/repositories"
	"trinity/internal/router"
	"trinity/internal/service"
	"trinity/pkg/database"
	"trinity/pkg/validator"

	"github.com/gin-gonic/gin"
)

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

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Initialize the database
	database.InitializeDB(cfg.DataBaseConfig.URL)
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	defer database.CloseDB()
	database.Migrate()
	// Initialize the data validator
	validator.InitValidator()
	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)
	campaignRepo := repositories.NewCampaignRepository(database.DB)
	voucherRepo := repositories.NewVoucherRepository(database.DB)
	voucherUserRepo := repositories.NewVoucherUserRepository(database.DB)

	userService := service.NewUserService(userRepo, campaignRepo, voucherRepo, voucherUserRepo)
	authService := service.NewAuthService(userRepo)
	voucherService := service.NewVoucherService(voucherRepo)
	campaignService := service.NewCampaignService(campaignRepo, voucherRepo)

	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)
	voucherHandler := handler.NewVoucherHandler(voucherService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// Setup routes
	router := router.SetupRouter(userHandler, authHandler, voucherHandler, campaignHandler, cfg)
	router.Run(cfg.HTTPServer.Port)
}
