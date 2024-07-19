package services

import (
	"database/sql"
)

type Services struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Services {
	return &Services{
		db: db,
	}
}
