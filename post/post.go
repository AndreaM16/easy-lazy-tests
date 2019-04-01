package post

import (
	"github.com/google/uuid"
)

// Post represents a post.
type Post struct {
	ID      uuid.UUID `json:"ID"`
	Content string    `json:"content"`
}

// Valid returns true if the post is valid.
func (p Post) Valid() bool {
	return p.ID == uuid.Nil && p.Content != ""
}
