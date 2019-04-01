package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andream16/personal-go-projects/posts/internal/serializer"
	mockserializer "github.com/andream16/personal-go-projects/posts/internal/serializer/mock"
	"github.com/andream16/personal-go-projects/posts/posts"
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

		const content = "Today I feel ok"

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
				"content" : "`+content+`"
			}`),
		)
		rr := httptest.NewRecorder()

		mockService.EXPECT().Add(posts.Post{
			Content: content,
		}).Return(nil, errors.New("someError"))

		h.add(req, rr)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}

	})

	t.Run("should return 500 because the service failed for adding a new post", func(t *testing.T) {

		const content = "Today I feel ok"

		post := posts.Post{
			Content: content,
		}
		newPost := &posts.Post{
			ID:      uuid.New(),
			Content: content,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)
		mockSerializer := mockserializer.NewMockSerializer(ctrl)

		h := New(
			mockSerializer,
			mockService,
		)

		req := httptest.NewRequest(
			http.MethodPost,
			"/add",
			bytes.NewBufferString(`{
				"content" : "`+content+`"
			}`),
		)
		rr := httptest.NewRecorder()

		mockSerializer.EXPECT().Deserialize(gomock.Any(), gomock.Any()).SetArg(1, post).Return(nil)
		mockService.EXPECT().Add(posts.Post{
			Content: content,
		}).Return(newPost, nil)
		mockSerializer.EXPECT().Serialize(newPost).Return(nil, errors.New("some error"))

		h.add(req, rr)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}

	})

	t.Run("should return 201", func(t *testing.T) {

		const content = "Today I feel ok"
		var (
			ID   = uuid.New()
			post = posts.Post{
				ID:      ID,
				Content: content,
			}
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
				"content" : "`+content+`"
			}`),
		)
		rr := httptest.NewRecorder()

		mockService.EXPECT().Add(posts.Post{
			Content: content,
		}).Return(&posts.Post{
			ID:      ID,
			Content: content,
		}, nil)

		h.add(req, rr)

		if rr.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d", rr.Code)
		}

		var p posts.Post

		err := serializer.Deserialize(rr.Body, &p)
		if err != nil {
			t.Fatal(err)
		}

		if post.ID != p.ID {
			t.Fatalf("expected ID %v, got %v", post.ID, p.ID)
		}

		if post.Content != p.Content {
			t.Fatalf("expected Content %v, got %v", post.ID, p.ID)
		}

	})

}
