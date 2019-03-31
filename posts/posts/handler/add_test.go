package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andream16/personal-go-projects/posts/internal/serializer"
	"github.com/andream16/personal-go-projects/posts/posts"
	"github.com/andream16/personal-go-projects/posts/posts/service"
	mockservice "github.com/andream16/personal-go-projects/posts/posts/service/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestHandler_add(t *testing.T) {

	t.Run("should return 400 because request's body is empty", func(t *testing.T) {

		serializer := serializer.HTTP{}

		h := New(
			serializer,
			nil,
		)

		req := httptest.NewRequest(
			http.MethodPost,
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
			serializer,
			nil,
		)

		req := httptest.NewRequest(
			http.MethodPost,
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
			ID      = "fe07cfbe-2e44-4c4c-8c81-ec7dfbf84510"
			content = "Today I feel ok"
		)

		serializer := serializer.HTTP{}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)

		h := New(
			serializer,
			mockService,
		)

		req := httptest.NewRequest(
			http.MethodPost,
			"/add",
			bytes.NewBufferString(`{
				"ID" : "`+ID+`",
				"content" : "`+content+`"
			}`),
		)
		rr := httptest.NewRecorder()

		mockService.EXPECT().Add(posts.Post{
			ID:      uuid.MustParse(ID),
			Content: content,
		}).Return(errors.New("someError"))

		h.add(req, rr)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}

	})

	t.Run("should return 409 because post already exists", func(t *testing.T) {

		const (
			ID      = "fe07cfbe-2e44-4c4c-8c81-ec7dfbf84510"
			content = "Today I feel ok"
		)

		serializer := serializer.HTTP{}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)

		h := New(
			serializer,
			mockService,
		)

		req := httptest.NewRequest(
			http.MethodPost,
			"/add",
			bytes.NewBufferString(`{
				"ID" : "`+ID+`",
				"content" : "`+content+`"
			}`),
		)
		rr := httptest.NewRecorder()

		mockService.EXPECT().Add(posts.Post{
			ID:      uuid.MustParse(ID),
			Content: content,
		}).Return(service.ErrAlreadyExists)

		h.add(req, rr)

		if rr.Code != http.StatusConflict {
			t.Fatalf("expected 409, got %d", rr.Code)
		}

	})

	t.Run("should return 201", func(t *testing.T) {

		const (
			ID      = "fe07cfbe-2e44-4c4c-8c81-ec7dfbf84510"
			content = "Today I feel ok"
		)

		serializer := serializer.HTTP{}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)

		h := New(
			serializer,
			mockService,
		)

		req := httptest.NewRequest(
			http.MethodPost,
			"/add",
			bytes.NewBufferString(`{
				"ID" : "`+ID+`",
				"content" : "`+content+`"
			}`),
		)
		rr := httptest.NewRecorder()

		mockService.EXPECT().Add(posts.Post{
			ID:      uuid.MustParse(ID),
			Content: content,
		}).Return(nil)

		h.add(req, rr)

		if rr.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d", rr.Code)
		}

	})

}
