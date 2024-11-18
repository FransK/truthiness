package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Collection interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{}
}
