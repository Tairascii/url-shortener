package handler

import (
	"context"
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
