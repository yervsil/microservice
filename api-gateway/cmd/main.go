package main

import (
	"log"

	"github.com/yervsil/api_gateway/config"
	"github.com/yervsil/api_gateway/internal/clients"
	"github.com/yervsil/api_gateway/internal/handler"
	"github.com/yervsil/api_gateway/internal/server"
	"github.com/yervsil/api_gateway/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	appLogger := logger.New(cfg.Env)
	appLogger.Info("Starting api-gateway server")

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		appLogger.Fatal(err)
	}
	userClient := clients.InitUserClient(conn)
	clients := clients.NewApiClients(userClient)
	handler := handler.NewHandler(appLogger, clients)

	srv := server.NewServer(cfg, handler.Routes())
	
	if err := srv.Run(); err != nil {
		appLogger.Fatal(err)
	}
}