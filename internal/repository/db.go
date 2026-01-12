package repository

import (
	"fmt"
	"log"

	"uam-golang/internal/config"
	"uam-golang/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database successfully")

	// Auto Migrate
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
	}
}
