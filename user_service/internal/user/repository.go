package user

import (
	"context"

	"github.com/yervsil/user_service/internal/user/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user *domain.User) (string, error)
	GetByCredentials(ctx context.Context, email, password string) (*domain.User, error)
}