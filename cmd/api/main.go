package main

import (
	"log"
	"net/http"

	"uam-golang/internal/config"
	"uam-golang/internal/handlers"
	"uam-golang/internal/middleware"
	"uam-golang/internal/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "uam-golang/docs" // Import generated docs
)

// @title           UAM User Access Management API
// @version         1.0
// @description     A secure, production-ready REST API for user authentication and management.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Connect to Database
	repository.ConnectDB(cfg)

	// 3. Setup Gin Router
	r := gin.Default()

	// 4. Routes
	api := r.Group("/api/v1")
	{
		// Swagger
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Health
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// Auth
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", handlers.Register)
			authGroup.POST("/login", handlers.Login)
		}

		// Users (Protected)
		userGroup := api.Group("/users")
		userGroup.Use(middleware.AuthMiddleware())
		{
			userGroup.GET("/me", handlers.GetProfile)
			userGroup.GET("", handlers.GetAllUsers)
		}
	}

	// 5. Start Server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
