package handler

import (
	"net/http"

	"github.com/andream16/easy-lazy-tests/post/service"

	"github.com/pkg/errors"
)

func (h *Handler) find(r *http.Request, w http.ResponseWriter) {

	ID := r.URL.Query().Get("ID")

	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, err := h.service.Find(ID)
	if err != nil {
		switch errors.Cause(err) {
		case service.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	b, err := h.serializer.Serialize(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)

}
