package repository

import (
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsernameOrEmail(username, email string) (*models.User, error)
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindByUsernameOrEmail(username, email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ? OR email = ?", username, email).First(&user).Error
	return &user, err
}

func (r *authRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
