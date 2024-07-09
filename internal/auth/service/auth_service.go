package service

import (
	"errors"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/repository"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/pkg/jwt"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/pkg/validators"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(authInput models.RegisterAuthInput) (*models.User, error)
	LoginUser(authInput models.LoginAuthInput) (string, string, error)
	GetUserByEmail(email string) (*models.User, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) RegisterUser(authInput models.RegisterAuthInput) (*models.User, error) {
	if err := validators.ValidatePassword(authInput.Password); err != nil {
		return nil, err
	}

	role, err := validators.ValidateUserRole(string(authInput.Role))
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  authInput.Username,
		Password:  string(passwordHash),
		Email:     authInput.Email,
		FirstName: authInput.FirstName,
		LastName:  authInput.LastName,
		Role:      role,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) LoginUser(authInput models.LoginAuthInput) (string, string, error) {
	user, err := s.repo.FindByEmail(authInput.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authInput.Password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, refreshToken, err := jwt.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}
