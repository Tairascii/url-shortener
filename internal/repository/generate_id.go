package repository

import (
	"context"
	"errors"

	query "github.com/Tairascii/url-shortener/db/query/repository"
	"github.com/Tairascii/url-shortener/internal/model"
)

func (r *Repo) GenerateID(ctx context.Context) (model.ID, error) {
	q := query.New(r.db)
	id, err := q.GenerateID(ctx, &query.GenerateIDParams{
		Column1: r.datacenterID,
		Column2: r.nodeID,
	})
	if err != nil {
		return 0, errors.Join(ErrOnQuery, err)
	}

	return id, nil
}
