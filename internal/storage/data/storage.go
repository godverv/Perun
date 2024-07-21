package data

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/storage/data/connections_cache"
	"github.com/Red-Sock/Perun/internal/storage/data/deploy_log"
	"github.com/Red-Sock/Perun/internal/storage/data/deploy_templates"
	"github.com/Red-Sock/Perun/internal/storage/data/instances"
	"github.com/Red-Sock/Perun/internal/storage/data/nodes"
	"github.com/Red-Sock/Perun/internal/storage/data/resources"
	"github.com/Red-Sock/Perun/internal/storage/data/services"
)

const (
	inMemory = "file::memory:?mode=memory&cache=shared"
	inFile   = "./data/sqlite-database.db"
)

type Store struct {
	nodes           *nodes.Provider
	services        *services.Provider
	resources       *resources.Provider
	instances       *instances.Provider
	deployTemplates *deploy_templates.Provider

	deployLog *deploy_log.Provider

	connectionsCache *connections_cache.ConnectionCache
}

func NewStorage(conn *sql.DB) storage.Data {
	return &Store{
		nodes:           nodes.NewStorage(conn),
		services:        services.NewStorage(conn),
		resources:       resources.New(conn),
		instances:       instances.New(conn),
		deployTemplates: deploy_templates.New(conn),

		deployLog: deploy_log.New(conn),

		connectionsCache: connections_cache.NewConnectionCache(),
	}
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

func (s *Store) Instances() storage.Instances {
	return s.instances
}

func (s *Store) DeployTemplates() storage.ResourceConstructors {
	return s.deployTemplates
}

func (s *Store) Connections() storage.ConnectionCache {
	return s.connectionsCache
}

func (s *Store) DeployLogs() storage.DeployLogs {
	return s.deployLog
}
