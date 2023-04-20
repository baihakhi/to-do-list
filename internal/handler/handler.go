package handler

import (
	"github.com/baihakhi/to-do-list/internal/service"
)

type Handler struct {
	service service.Service
}

func InitiHandler(services service.Service) *Handler {
	return &Handler{
		service: services,
	}
}
