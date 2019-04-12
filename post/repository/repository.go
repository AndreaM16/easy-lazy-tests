//go:generate mockgen -self_package github.com/andream16/easy-lazy-tests/posts/repository -source=repository.go -destination mock/repository_mock.go

package repository

import (
	"github.com/andream16/easy-lazy-tests/post"

	"github.com/google/uuid"
)

// Repositorer represents the repository interface.
type Repositorer interface {
	Insert(post.Post) error
	Find(uuid.UUID) (*post.Post, error)
}
