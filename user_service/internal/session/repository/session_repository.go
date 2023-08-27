package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/yervsil/user_service/config"
	"github.com/yervsil/user_service/internal/session"
	"github.com/yervsil/user_service/internal/user/domain"
)

const (
	basePrefix = "sessions:"
)
                                                                                         
// Session repository
type sessionRepo struct {
	redisClient *redis.Client
	basePrefix  string
	cfg         *config.Config
}

// Session repository constructor
func NewSessionRepository(redisClient *redis.Client, cfg *config.Config) session.SessRepository {
	return &sessionRepo{redisClient: redisClient, basePrefix: basePrefix, cfg: cfg}
}

func(sr *sessionRepo) CreateSession(ctx context.Context, session *domain.Session, expire int) (string, error) {

	session.SessionID = uuid.New().String()
	sessionKey := sr.createKey(session.SessionID)

	sessBytes, err := json.Marshal(&session)
	if err != nil {
		return "", err
	}
	if err = sr.redisClient.Set(ctx, sessionKey, sessBytes, time.Second*time.Duration(expire)).Err(); err != nil {
		return "", err
	}
	return session.SessionID, nil
}

func (s *sessionRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.basePrefix, sessionID)
}