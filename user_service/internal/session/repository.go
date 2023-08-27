package session

import (
	"context"

	"github.com/yervsil/user_service/internal/user/domain"
)

type SessRepository interface {
	CreateSession(ctx context.Context, session *domain.Session, expire int) (string, error)
	// GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error)
	// DeleteByID(ctx context.Context, sessionID string) error
}
