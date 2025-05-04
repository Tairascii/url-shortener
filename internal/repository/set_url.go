package repository

import (
	"context"
	"errors"

	query "github.com/Tairascii/url-shortener/db/query/repository"
	"github.com/Tairascii/url-shortener/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repo) SetURL(ctx context.Context, id model.ID, longURL, shortURL string) error {
	q := query.New(r.db)
	err := q.SetUrl(ctx, &query.SetUrlParams{
		LongUrl:  pgtype.Text{String: longURL, Valid: true},
		ShortUrl: pgtype.Text{String: shortURL, Valid: true},
		ID:       id,
	})
	if err != nil {
		return errors.Join(ErrOnQuery, err)
	}

	return nil
}
