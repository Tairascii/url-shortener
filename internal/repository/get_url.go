package repository

import (
	"context"
	"errors"

	query "github.com/Tairascii/url-shortener/db/query/repository"
	"github.com/Tairascii/url-shortener/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repo) GetURL(ctx context.Context, shortURL string) (string, error) {
	q := query.New(r.db)
	url, err := q.GetFullUrl(ctx, pgtype.Text{String: shortURL, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", model.ErrURLNotFound
		}
		return "", errors.Join(ErrOnQuery, err)
	}

	return url.String, nil
}
