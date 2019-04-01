package handler

import (
	"github.com/andream16/easy-lazy-tests/internal/serializer"
	"github.com/andream16/easy-lazy-tests/post/service"
)

// Handler represents an handler.
type Handler struct {
	serializer serializer.Serializer
	service    service.Servicer
}

// New returns a new handler.
func New(
	serializer serializer.Serializer,
	service service.Servicer,
) *Handler {
	return &Handler{
		serializer: serializer,
		service:    service,
	}
}
