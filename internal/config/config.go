package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost        string        `mapstructure:"DB_HOST"`
	DBUser        string        `mapstructure:"DB_USER"`
	DBPassword    string        `mapstructure:"DB_PASSWORD"`
	DBName        string        `mapstructure:"DB_NAME"`
	DBPort        string        `mapstructure:"DB_PORT"`
	JWTSecret     string        `mapstructure:"JWT_SECRET"`
	JWTExpiration time.Duration `mapstructure:"JWT_EXPIRATION"`
	ServerPort    string        `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("JWT_EXPIRATION", "24h")
	viper.SetDefault("DB_PORT", "5432")

	// Read from .env file if it exists
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // Read from environment variables

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config file: %s", err)
			return
		}
		// Config file not found is fine if we're using env vars
		log.Println("No .env file found, using environment variables")
	}

	err = viper.Unmarshal(&config)
	return
}
