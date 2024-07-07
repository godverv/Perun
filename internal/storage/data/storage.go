package data

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/storage/data/connections_cache"
	"github.com/Red-Sock/Perun/internal/storage/data/nodes"
	"github.com/Red-Sock/Perun/internal/storage/data/resources"
)

const (
	inMemory = "file::memory:?mode=memory&cache=shared"
	inFile   = "./data/sqlite-database.db"
)

type Store struct {
	nodes            *nodes.Provider
	resources        *resources.Storage
	connectionsCache *connections_cache.ConnectionCache
}

func NewStorage(conn *sql.DB) (storage.Data, error) {
	return &Store{
		nodes:     nodes.NewStorage(conn),
		resources: resources.NewProvider(conn),

		connectionsCache: connections_cache.NewConnectionCache(),
	}, nil
}

func (s *Store) Nodes() storage.Nodes {
	return s.nodes
}

func (s *Store) Resources() storage.Resources {
	return s.resources
}
func (s *Store) Connections() storage.ConnectionCache {
	return s.connectionsCache
}
