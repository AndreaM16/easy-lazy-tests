package post

import (
	"testing"

	"github.com/google/uuid"
)

func TestPost_Valid(t *testing.T) {

	testCases := []struct {
		post     Post
		expected bool
	}{
		{
			post:     Post{},
			expected: false,
		},
		{
			post: Post{
				ID: uuid.New(),
			},
			expected: false,
		},
		{
			post: Post{
				Content: "some content",
			},
			expected: true,
		},
	}

	for _, testCase := range testCases {

		valid := testCase.post.Valid()

		if testCase.expected != valid {
			t.Fatalf("expected %t, got %t", testCase.expected, valid)
		}

	}

}
