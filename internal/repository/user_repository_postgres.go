package repository

import (
	"github.com/pujilesmana/chat-app/internal/domain"
	"gorm.io/gorm"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepositoryPostgres(db *gorm.DB) domain.UserRepository {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryPostgres) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPostgres) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
