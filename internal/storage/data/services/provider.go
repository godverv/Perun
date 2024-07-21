package services

import (
	"database/sql"
)

type Provider struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Provider {
	return &Provider{
		db: db,
	}
}
