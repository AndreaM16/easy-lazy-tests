//go:generate mockgen -self_package github.com/andream16/personal-go-projects/posts/posts/service -source=service.go -destination mock/service_mock.go

package service

import (
	"github.com/andream16/personal-go-projects/posts/posts"
	"github.com/andream16/personal-go-projects/posts/posts/repository"

	"github.com/pkg/errors"
)

// Service related errors.
var (
	ErrAlreadyExists = errors.New("post_already_exists")
	ErrNotFound      = errors.New("post_not_found")
)

// Servicer is the service interface.
type Servicer interface {
	Add(posts.Post) error
	Find(string) (*posts.Post, error)
}

// Service represents the service.
type Service struct {
	repository repository.Repositorer
}

// New returns a new service.
func New(repository repository.Repositorer) *Service {
	return &Service{
		repository: repository,
	}
}