package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrOnQuery = errors.New("query error")
)

type Repo struct {
	db           *pgxpool.Pool
	datacenterID uint8
	shardID      uint8
}

func New(db *pgxpool.Pool, datacenterID, shardID uint8) *Repo {
	return &Repo{
		db:           db,
		datacenterID: datacenterID,
		shardID:      shardID,
	}
}
