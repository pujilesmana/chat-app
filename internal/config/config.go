package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	JWTSecret  string
	BaseURL    string
	AppPort    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		BaseURL:    os.Getenv("BASE_URL"),
		AppPort:    os.Getenv("APP_PORT"),
	}
}
