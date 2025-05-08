package handler

import (
	"context"
	"net/http"
)

const (
	shortURLPathParam = "short_url"
)

type UseCase interface {
	AddURL(ctx context.Context, url string) (string, error)
	GetURL(ctx context.Context, shortURL string) (string, error)
}

type Handler struct {
	useCase UseCase
}

func New(useCase UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func AttachRoutes(h *Handler) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /api/url/add", h.AddURL)
	router.HandleFunc("GET /api/url/{short_url}", h.GetURL)
	return router
}

func UrlFromPath(r *http.Request) string {
	shortURL := r.PathValue(shortURLPathParam)
	return shortURL
}
