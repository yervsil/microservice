package usecase

import (
	"context"

	"github.com/yervsil/user_service/internal/session"
	"github.com/yervsil/user_service/internal/user/domain"
)

type sessionUscase struct {
	sessionRepo session.SessRepository
}

func NewSessionUseCase(sessionRepo session.SessRepository) *sessionUscase {
	return &sessionUscase{sessionRepo: sessionRepo}
}

// Create new session
func (u *sessionUscase) CreateSession(ctx context.Context, session *domain.Session, expire int) (string, error) {
	
	return u.sessionRepo.CreateSession(ctx, session, expire)
}