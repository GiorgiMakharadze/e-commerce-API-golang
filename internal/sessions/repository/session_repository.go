package repository

import (
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
	"gorm.io/gorm"
)

type SessionRepository interface {
	CreateSession(session *models.Session) error
	GetSessionByToken(token string) (*models.Session, error)
	DeleteSessionByToken(token string) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db}
}

func (r *sessionRepository) CreateSession(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepository) GetSessionByToken(token string) (*models.Session, error) {
	var session models.Session
	err := r.db.Where("access_token = ? OR refresh_token = ?", token, token).First(&session).Error
	return &session, err
}

func (r *sessionRepository) DeleteSessionByToken(token string) error {
	return r.db.Where("access_token = ? OR refresh_token = ?", token, token).Delete(&models.Session{}).Error
}
