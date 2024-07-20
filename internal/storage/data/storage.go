package data

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/storage/data/connections_cache"
	"github.com/Red-Sock/Perun/internal/storage/data/deploy_log"
	"github.com/Red-Sock/Perun/internal/storage/data/nodes"
	"github.com/Red-Sock/Perun/internal/storage/data/resources"
	"github.com/Red-Sock/Perun/internal/storage/data/services"
)

const (
	inMemory = "file::memory:?mode=memory&cache=shared"
	inFile   = "./data/sqlite-database.db"
)

type Store struct {
	nodes     *nodes.Provider
	services  *services.Services
	resources *resources.Provider
	deployLog *deploy_log.Provider

	connectionsCache *connections_cache.ConnectionCache
}

func NewStorage(conn *sql.DB) (storage.Data, error) {
	return &Store{
		nodes:     nodes.NewStorage(conn),
		services:  services.NewStorage(conn),
		resources: resources.New(conn),
		deployLog: deploy_log.New(conn),

		connectionsCache: connections_cache.NewConnectionCache(),
	}, nil
}

func (s *Store) Resources() storage.Resources {
	return s.resources
}

func (s *Store) Nodes() storage.Nodes {
	return s.nodes
}

func (s *Store) Services() storage.Services {
	return s.services
}
func (s *Store) Connections() storage.ConnectionCache {
	return s.connectionsCache
}

func (s *Store) DeployLogs() storage.DeployLogs {
	return s.deployLog
}
