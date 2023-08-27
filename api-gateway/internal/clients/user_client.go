package clients

import (
	"context"

	"github.com/yervsil/api_gateway/internal/dto"
	"github.com/yervsil/api_gateway/proto/user"
	"google.golang.org/grpc"
)

type UserClient struct {
	uc user.UserServiceClient
}

func InitUserClient(conn *grpc.ClientConn) UserClient {
	uc := user.NewUserServiceClient(conn)
	return UserClient{
		uc:  uc,
	}
}

func (gc *UserClient) CreateUser(ctx context.Context, dto dto.SignUp) (string, error){
	request := user.SignUpRequest{Username: dto.Username, Email: dto.Email, Password: dto.Password}

	response, err := gc.uc.SignUp(context.Background(), &request)
	if err != nil {
		return "", err
	}

	return response.GetUserID(), nil 
}

func (gc *UserClient) Login(ctx context.Context, dto dto.SignIn) (string, error){
	request := user.SignInRequest{Email: dto.Email, Password: dto.Password}

	response, err := gc.uc.SignIn(context.Background(), &request)
	if err != nil {
		return "", err
	}

	return response.GetSessionId(), nil 
}