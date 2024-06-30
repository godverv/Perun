package nodes

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Provider struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Provider {
	return &Provider{
		db: db,
	}
}
