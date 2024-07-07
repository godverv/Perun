package resources

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func NewProvider(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}
