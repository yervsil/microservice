package handler

import (
	"github.com/yervsil/api_gateway/internal/clients"
	"github.com/yervsil/api_gateway/pkg/logger"
)

type Handler struct {
	AppLogger *logger.Logger
	Clients *clients.Clients
}

func NewHandler(AppLogger *logger.Logger, Clients *clients.Clients) *Handler {
	return &Handler{AppLogger: AppLogger, Clients: Clients}
}