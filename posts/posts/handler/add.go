package handler

import (
	"net/http"

	"github.com/andream16/personal-go-projects/posts/posts"
	"github.com/andream16/personal-go-projects/posts/posts/service"

	"github.com/pkg/errors"
)

func (h *Handler) add(r *http.Request, w http.ResponseWriter) {

	var post posts.Post

	if error := h.serializer.Deserialize(r.Body, &post); error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !post.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.service.Add(post)

	if err != nil {
		switch errors.Cause(err) {
		case service.ErrAlreadyExists:
			w.WriteHeader(http.StatusConflict)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)

}
