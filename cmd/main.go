package main

import (
	"fmt"
	"github.com/pujilesmana/chat-app/internal/config"
	httpDelivery "github.com/pujilesmana/chat-app/internal/delivery/http"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=chatapp sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// Set up your HTTP routes
	http.HandleFunc("/register", httpDelivery.RegisterHandler(db, cfg.JWTSecret))
	http.HandleFunc("/login", httpDelivery.LoginHandler(db, cfg.JWTSecret))

	// Start the HTTP server
	addr := fmt.Sprintf("%s:%s", cfg.BaseURL, cfg.AppPort)
	log.Info().Msgf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
