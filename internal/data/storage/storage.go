package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/data/storage/nodes"
	"github.com/Red-Sock/Perun/internal/data/storage/resources"
)

const (
	inMemory = "file::memory:?mode=memory&cache=shared"
	inFile   = "./data/sqlite-database.db"
)

type Store struct {
	nodes     *nodes.Provider
	resources *resources.Storage
}

func NewStorage(conn *sql.DB) (data.Storage, error) {
	return &Store{
		nodes:     nodes.NewStorage(conn),
		resources: resources.NewProvider(conn),
	}, nil
}

func (s *Store) Nodes() data.Nodes {
	return s.nodes
}

func (s *Store) Resources() data.Resources {
	return s.resources
}
