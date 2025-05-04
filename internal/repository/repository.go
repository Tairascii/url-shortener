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
	datacenterID int64
	nodeID       int64
}

func New(db *pgxpool.Pool, datacenterID, nodeID int64) *Repo {
	return &Repo{
		db:           db,
		datacenterID: datacenterID,
		nodeID:       nodeID,
	}
}
