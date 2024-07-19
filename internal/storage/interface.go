package storage

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

type Data interface {
	Nodes() Nodes
	Services() Services

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

	Get(ctx context.Context, name string) (*domain.Service, error)
	List(ctx context.Context, serviceNamePattern string) ([]domain.Service, error)
	ListChildren(ctx context.Context, parentName string) ([]domain.Service, error)
}

type ConnectionCache interface {
	Add(nodes ...domain.Node)
	Get(names ...string) ([]domain.Node, error)
}

type Row interface {
	Scan(dest ...interface{}) error
}
