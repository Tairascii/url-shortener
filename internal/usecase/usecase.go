package usecase

import (
	"context"

	"github.com/Tairascii/url-shortener/pkg"
)

type Repo interface {
	GenerateID(ctx context.Context) (int64, error)
	SetURL(ctx context.Context, id int64, longURL, shortURL string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
}

type UseCase struct {
	repo Repo
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) AddURL(ctx context.Context, url string) (string, error) {
	// should be in transaction
	id, err := u.repo.GenerateID(ctx)
	if err != nil {
		return "", err
	}

	generatedURL := pkg.GenerateURL(id)
	err = u.repo.SetURL(ctx, id, url, generatedURL)
	if err != nil {
		return "", err
	}
	return generatedURL, nil
}

func (u *UseCase) GetURL(ctx context.Context, shortURL string) (string, error) {
	url, err := u.repo.GetURL(ctx, shortURL)
	if err != nil {
		return "", err
	}

	return url, nil
}
