package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tairascii/url-shortener/pkg/response_writers"
)

type AddURLPayload struct {
	Url string `json:"url"`
}

type AddURLResponse struct {
	ShortURL string `json:"short_url"`
}

func (h *Handler) AddURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	payload, err := validateAddURLPayload(r)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrl, err := h.useCase.AddURL(ctx, payload.Url)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response_writers.JSONResponseWriter(w, http.StatusOK, AddURLResponse{ShortURL: shortUrl})
}

func validateAddURLPayload(r *http.Request) (AddURLPayload, error) {
	var payload AddURLPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return AddURLPayload{}, err
	}
	return payload, nil
}
