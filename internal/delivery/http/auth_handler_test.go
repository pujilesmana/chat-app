package http

import (
	"bytes"
	"encoding/json"
	"github.com/pujilesmana/chat-app/internal/domain"
	"github.com/pujilesmana/chat-app/internal/repository"
	"github.com/pujilesmana/chat-app/internal/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=chatapp_test port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&domain.User{})
	return db
}

func TestRegisterHandler(t *testing.T) {
	db := setupTestDB()
	handler := RegisterHandler(db, "secret")

	reqBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"password": "password123",
	})

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Response should be 201 Created")
	assert.Contains(t, rr.Body.String(), "user registered", "Response body should contain success message")
}

func TestLoginHandler(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUserRepositoryPostgres(db)
	usecase := usecase.NewAuthUsecase(repo, "secret")
	handler := LoginHandler(db, "secret")

	// Register user
	usecase.Register("testuser", "password123")

	reqBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"password": "password123",
	})

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Response should be 200 OK")
	assert.Contains(t, rr.Body.String(), "token", "Response body should contain token")
}
