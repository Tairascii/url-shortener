package model

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
)

type ID = int64
