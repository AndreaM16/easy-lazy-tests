package handler

import (
	"github.com/andream16/personal-go-projects/posts/posts/service"
	"github.com/andream16/personal-go-projects/posts/internal/logger"
	"github.com/andream16/personal-go-projects/posts/internal/serializer"
)

// Handler represents an handler.
type Handler struct {
	logger logger.Logger
	serializer serializer.Serializer
	service service.Servicer
}

// New returns a new handler.
func New(
	logger logger.Logger,
	serializer serializer.Serializer, 
	service service.Servicer,
) *Handler {
	return &Handler{
		logger: logger,
		serializer: serializer,
		service: service,
	}
}