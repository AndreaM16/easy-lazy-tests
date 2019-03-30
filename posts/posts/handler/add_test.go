package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andream16/personal-go-projects/posts/internal/serializer"
	"github.com/andream16/personal-go-projects/posts/posts"
	mockservice "github.com/andream16/personal-go-projects/posts/posts/service/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestHandler_add(t *testing.T) {

	t.Run("should return 400 because request's body is empty", func(t *testing.T) {

		serializer := serializer.HTTP{}

		h := New(
			nil,
			serializer,
			nil,
		)

		req := httptest.NewRequest(
			http.MethodGet,
			"/add",
			nil,
		)
		rr := httptest.NewRecorder()

		h.add(req, rr)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}

	})

	t.Run("should return 400 because the request is invalid", func(t *testing.T) {

		serializer := serializer.HTTP{}

		h := New(
			nil,
			serializer,
			nil,
		)

		req := httptest.NewRequest(
			http.MethodGet,
			"/add",
			bytes.NewBufferString(`{}`),
		)
		rr := httptest.NewRecorder()

		h.add(req, rr)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}

	})

	t.Run("should return 500 because the service failed for adding a new post", func(t *testing.T) {

		const (
			ID      = "a845975d-d584-4e78-8920-17ea70f32b5f"
			content = "Today I feel ok"
		)

		serializer := serializer.HTTP{}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)

		h := New(
			nil,
			serializer,
			mockService,
		)

		mockService.EXPECT().Add(posts.Post{
			ID:      uuid.MustParse(ID),
			Content: content,
		}).Return(errors.New("someError"))

		req := httptest.NewRequest(
			http.MethodGet,
			"/add",
			bytes.NewBufferString(`{
				"ID" : `+ID+`,
				"content" : `+content+`
			}`),
		)
		rr := httptest.NewRecorder()

		h.add(req, rr)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}

	})

}
