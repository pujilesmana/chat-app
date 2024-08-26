package usecase

import (
	"errors"
	"github.com/pujilesmana/chat-app/internal/domain"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepository struct {
	users map[string]*domain.User
}

func (m *mockUserRepository) Create(user *domain.User) error {
	if _, exists := m.users[user.Username]; exists {
		return errors.New("user already exists")
	}
	m.users[user.Username] = user
	return nil
}

func (m *mockUserRepository) GetByUsername(username string) (*domain.User, error) {
	user, exists := m.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *mockUserRepository) GetByID(id uint) (*domain.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func TestRegister(t *testing.T) {
	mockRepo := &mockUserRepository{users: make(map[string]*domain.User)}
	usecase := NewAuthUsecase(mockRepo, "secret")

	err := usecase.Register("testuser", "password123")
	assert.NoError(t, err, "Register should not return an error")

	user, err := mockRepo.GetByUsername("testuser")
	assert.NoError(t, err, "User should be found")
	assert.NotNil(t, user, "User should be created")
	assert.Equal(t, "testuser", user.Username, "Username should match")
}

func TestLogin(t *testing.T) {
	mockRepo := &mockUserRepository{users: make(map[string]*domain.User)}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockRepo.Create(&domain.User{
		ID:        1,
		Username:  "testuser",
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	usecase := NewAuthUsecase(mockRepo, "secret")

	token, err := usecase.Login("testuser", "password123")
	assert.NoError(t, err, "Login should not return an error")
	assert.NotEmpty(t, token, "Token should be returned")

	// Validate the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	assert.NoError(t, err, "Token should be valid")
	assert.True(t, parsedToken.Valid, "Token should be valid")
}
