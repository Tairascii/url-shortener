package handler

import (
	"errors"
	"net/http"

	"github.com/Tairascii/url-shortener/internal/model"
	"github.com/Tairascii/url-shortener/pkg/response_writers"
)

type GetURLResponse struct {
	FullURL string `json:"full_url"`
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	shortURL := UrlFromPath(r)

	fullURL, err := h.useCase.GetURL(ctx, shortURL)
	if err != nil {
		if errors.Is(err, model.ErrURLNotFound) {
			response_writers.ErrorResponseWriter(w, model.ErrURLNotFound.Error(), http.StatusNotFound)
			return
		}
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response_writers.JSONResponseWriter(w, http.StatusOK, GetURLResponse{FullURL: fullURL})
}
