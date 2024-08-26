package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pujilesmana/chat-app/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthUsecase struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

func NewAuthUsecase(userRepo domain.UserRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *AuthUsecase) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username:  username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u.userRepo.Create(user)
}

func (u *AuthUsecase) Login(username, password string) (string, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
