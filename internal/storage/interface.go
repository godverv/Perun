package storage

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

type Row interface {
	Scan(dest ...interface{}) error
}

type Data interface {
	Nodes() Nodes
	Services() Services
	Resources() Resources

	Connections() ConnectionCache
}

type Nodes interface {
	GetConnection(ctx context.Context, key string) (*domain.VelezConnection, error)
	SaveConnection(ctx context.Context, in domain.VelezConnection) error

	ListNodes(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConnection, error)
	ListConnections(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConnection, error)
	ListLeastUsedNodes(ctx context.Context, req domain.PickNodesReq) ([]domain.VelezConnection, error)
}

type Services interface {
	Upsert(ctx context.Context, resource ...domain.Service) error
	UpdateState(ctx context.Context, resource domain.Service) error

	Get(ctx context.Context, name string) (domain.Service, error)
	List(ctx context.Context, serviceNamePattern string) ([]domain.Service, error)
}

type Resources interface {
	ListForService(ctx context.Context, name string) ([]domain.Resource, error)
	Upsert(ctx context.Context, deps ...domain.Resource) error
}

type ConnectionCache interface {
	Add(nodes ...domain.Node)
	Get(names ...string) ([]domain.Node, error)
}
