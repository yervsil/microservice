package user

import (
	"context"

	"github.com/yervsil/user_service/internal/user/domain"
)

// UseCase Product
type Usecase interface {
	CreateUser(ctx context.Context, user *domain.User) (string, error)
	Login(ctx context.Context, email, password string) ( *domain.User, error)
}