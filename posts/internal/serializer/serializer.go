//go:generate mockgen -self_package github.com/andream16/personal-go-projects/posts/internal/serializer -source=serializer.go -destination mock/serializer_mock.go

package serializer

import (
	"encoding/json"
	"io"
)

// Serializer represents the serializer interface.
type Serializer interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize(io.Reader, interface{}) error
}

// HTTP represents the concrete implementation of serializer for http.
type HTTP struct{}

// Serialize serializes the input.
func (HTTP) Serialize(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

// Deserialize deserializes the input into a given reference.
func (HTTP) Deserialize(r io.Reader, out interface{}) error {
	return json.NewDecoder(r).Decode(out)
}
