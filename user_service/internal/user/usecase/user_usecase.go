package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	"github.com/yervsil/user_service/internal/user/domain"
	"github.com/yervsil/user_service/internal/user"
)

const (
	salt = "felkwemfwlke42t23r23fsd"
)

type UserUsecase struct {
	userMongoRepo user.Repository
}

func NewUserUsecase(userMongoRepo user.Repository) *UserUsecase {
	return &UserUsecase{userMongoRepo: userMongoRepo}
}

func(uu *UserUsecase) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	userService := domain.User{Username: user.Username, Email: user.Email, Password: hashPasswordWithSalt(user.Password)}
	return uu.userMongoRepo.CreateUser(ctx, &userService)
}

func(uu *UserUsecase) Login(ctx context.Context, email, password string) ( *domain.User, error) {
	foundUser, err := uu.userMongoRepo.GetByCredentials(ctx, email, hashPasswordWithSalt(password))
	if err != nil {
		return nil, err
	}

	return foundUser, err
}



func hashPasswordWithSalt(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password + salt))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// func (s *UsersService) createSession(userId string) (Tokens, error) {
// 	var (
// 		res Tokens
// 		err error
// 	)

// 	res.AccessToken, err = s.tokenManager.NewJWT(userId, "user", s.accessTokenTTL)
// 	if err != nil {
// 		return res, err
// 	}

// 	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
// 	if err != nil {
// 		return res, err
// 	}

// 	session := domain.Session{
// 		RefreshToken: res.RefreshToken,
// 		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
// 	}

// 	err = s.repo.SetSession(userId, session)

// 	return res, err
// }