package server

import (
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/yervsil/user_service/config"
	"github.com/yervsil/user_service/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"

	//"github.com/yervsil/user_service/internal/user"
	sessionUsecase "github.com/yervsil/user_service/internal/session/usecase"
	userUsecase "github.com/yervsil/user_service/internal/user/usecase"
	userService "github.com/yervsil/user_service/internal/user/delivery/grpc"
	userRepository "github.com/yervsil/user_service/internal/user/repository"
	sessionRepository "github.com/yervsil/user_service/internal/session/repository"
	"github.com/yervsil/user_service/proto"
	"google.golang.org/grpc"
)

type server struct {
	log     *logger.Logger
	cfg     *config.Config
	mongoDB *mongo.Database
	redisClient *redis.Client
}

// NewServer constructor
func NewServer(log *logger.Logger, cfg *config.Config, mongoDB *mongo.Database, redisClient *redis.Client) *server {
	return &server{log: log, cfg: cfg, mongoDB: mongoDB, redisClient: redisClient}
}

// Run Start server
func (s *server) Run() error {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	userMongoRepo := userRepository.NewUserRepository(s.mongoDB)
	userRedisRepo := sessionRepository.NewSessionRepository(s.redisClient, s.cfg)
	userUC := userUsecase.NewUserUsecase(userMongoRepo)
	sesUC := sessionUsecase.NewSessionUseCase(userRedisRepo)

	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		s.log.Info(err.Error())
		return err
	}
	defer l.Close()

	
	grpcServer := grpc.NewServer()

	userService := userService.NewUserService(userUC, sesUC, s.log)
	proto.RegisterUserServiceServer(grpcServer, userService)
	
	if err := grpcServer.Serve(l); err != nil {
		s.log.Info(err.Error())
	}
	return nil
}