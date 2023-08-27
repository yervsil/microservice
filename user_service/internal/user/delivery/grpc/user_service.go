package grpc

import (
	"context"

	"github.com/yervsil/user_service/internal/session"
	"github.com/yervsil/user_service/internal/user"
	"github.com/yervsil/user_service/internal/user/domain"
	"github.com/yervsil/user_service/pkg/logger"
	"github.com/yervsil/user_service/pkg/utils"
	"github.com/yervsil/user_service/proto"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	userUC user.Usecase
	sessionUC session.SessUsecase
	logger *logger.Logger
}

func NewUserService(userUC user.Usecase, sessionUC session.SessUsecase, logger *logger.Logger) *UserService{
	return &UserService{userUC: userUC, sessionUC: sessionUC, logger: logger}
}

func (us *UserService) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	userDto := us.reqToDto(req)

	if err := utils.ValidateUser(userDto); err != nil {
		return nil, err
	}

	id, err := us.userUC.CreateUser(ctx, &userDto)
	if err != nil {
		return nil, err
	}

	return &proto.SignUpResponse{UserID: id}, nil
}

func (us *UserService) SignIn(ctx context.Context, r *proto.SignInRequest) (*proto.SignInResponse, error) {
	user, err := us.userUC.Login(ctx, r.GetEmail(), r.GetPassword())
	if err != nil {
		us.logger.Error("userUC.Login: %v", err)
		return nil, err
	}

	session, err := us.sessionUC.CreateSession(ctx, &domain.Session{
		UserID: user.ID,
	}, 300)
	if err != nil {
		us.logger.Error("sessUC.CreateSession: %v", err)
		return nil, err
	}

	return &proto.SignInResponse{User: us.userModelToProto(user), SessionId: session}, err
}

func (us *UserService) reqToDto(req *proto.SignUpRequest) domain.User {
	return domain.User{Username: req.GetUsername(), Email: req.GetEmail(), Password: req.GetPassword()}
}

func (us *UserService) userModelToProto(user *domain.User) *proto.User {
	userProto := &proto.User{
		Uuid:      user.ID.Hex(),
		Username: user.Username,
		Password: user.Password,
		Email: user.Email,
	}
	return userProto
}