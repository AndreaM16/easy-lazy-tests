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
