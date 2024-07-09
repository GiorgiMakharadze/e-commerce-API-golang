package service

import (
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/sessions/repository"
)

type SessionService interface {
	CreateSession(session *models.Session) error
	GetSessionByToken(token string) (*models.Session, error)
	DeleteSessionByToken(token string) error
}

type sessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) repository.SessionRepository {
	return &sessionService{repo}
}

func (s *sessionService) CreateSession(session *models.Session) error {
	return s.repo.CreateSession(session)
}

func (s *sessionService) GetSessionByToken(token string) (*models.Session, error) {
	return s.repo.GetSessionByToken(token)
}

func (s *sessionService) DeleteSessionByToken(token string) error {
	return s.repo.DeleteSessionByToken(token)
}
