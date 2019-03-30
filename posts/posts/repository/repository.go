package repository

import (
	"github.com/andream16/personal-go-projects/posts/posts"

	"github.com/google/uuid"
)

// Repositorer represents the repository interface.
type Repositorer interface {
	Insert(posts.Post) error
	Find(uuid.UUID) (*posts.Post, error)
}