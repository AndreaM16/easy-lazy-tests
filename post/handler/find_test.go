package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/andream16/easy-lazy-tests/internal/serializer"
	mockserializer "github.com/andream16/easy-lazy-tests/internal/serializer/mock"
	"github.com/andream16/easy-lazy-tests/post"
	"github.com/andream16/easy-lazy-tests/post/service"
	mockservice "github.com/andream16/easy-lazy-tests/post/service/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestHandler_find(t *testing.T) {

	t.Run("should return 400 because the post id is malformed", func(t *testing.T) {

		h := New(
			nil,
			nil,
		)

		req := httptest.NewRequest(
			http.MethodGet,
			"/",
			nil,
		)
		rr := httptest.NewRecorder()

		h.find(req, rr)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}

	})

	t.Run("should return 404 because no post is found for such ID", func(t *testing.T) {

		ID := uuid.New().String()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)

		h := New(
			nil,
			mockService,
		)

		req := httptest.NewRequest(
			http.MethodGet,
			"/{ID}",
			nil,
		)

		q := url.Values{}
		q.Add("ID", ID)
		req.URL.RawQuery = q.Encode()

		mockService.EXPECT().Find(ID).Return(nil, service.ErrNotFound)

		rr := httptest.NewRecorder()

		h.find(req, rr)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", rr.Code)
		}

	})

	t.Run("should return 500 because some error occurred during find post", func(t *testing.T) {

		ID := uuid.New().String()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)

		h := New(
			nil,
			mockService,
		)

		req := httptest.NewRequest(
			http.MethodGet,
			"/{ID}",
			nil,
		)

		q := url.Values{}
		q.Add("ID", ID)
		req.URL.RawQuery = q.Encode()

		mockService.EXPECT().Find(ID).Return(nil, errors.New("some error"))

		rr := httptest.NewRecorder()

		h.find(req, rr)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}

	})

	t.Run("should return 500 because some error occurred during serialization of a found post", func(t *testing.T) {

		ID := uuid.New().String()
		post := &post.Post{
			ID:      uuid.MustParse(ID),
			Content: "Meh, sounds good",
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
			http.MethodGet,
			"/{ID}",
			nil,
		)

		q := url.Values{}
		q.Add("ID", ID)
		req.URL.RawQuery = q.Encode()

		mockService.EXPECT().Find(ID).Return(post, nil)
		mockSerializer.EXPECT().Serialize(post).Return(nil, errors.New("some error"))

		rr := httptest.NewRecorder()

		h.find(req, rr)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected 500, got %d", rr.Code)
		}

	})

	t.Run("should return 200 for an existing post", func(t *testing.T) {

		var (
			ID           = uuid.New().String()
			expectedPost = &post.Post{
				ID:      uuid.MustParse(ID),
				Content: "Meh, sounds good",
			}
		)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockServicer(ctrl)
		serializer := serializer.HTTP{}

		h := New(
			serializer,
			mockService,
		)

		req := httptest.NewRequest(
			http.MethodGet,
			"/{ID}",
			nil,
		)

		q := url.Values{}
		q.Add("ID", ID)
		req.URL.RawQuery = q.Encode()

		mockService.EXPECT().Find(ID).Return(expectedPost, nil)

		rr := httptest.NewRecorder()

		h.find(req, rr)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var p post.Post

		err := serializer.Deserialize(rr.Body, &p)
		if err != nil {
			t.Fatal(err)
		}

		if expectedPost.ID != p.ID {
			t.Fatalf("expected ID %v, got %v", expectedPost.ID, p.ID)
		}

		if expectedPost.Content != p.Content {
			t.Fatalf("expected Content %v, got %v", expectedPost.ID, p.ID)
		}

	})

}
