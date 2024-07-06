package data

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

type Data interface {
	Nodes() Nodes
	Resources() Resources

	Connections() ConnectionCache
}

type Nodes interface {
	GetConnection(ctx context.Context, key string) (*domain.VelezConnection, error)
	SaveConnection(ctx context.Context, in domain.VelezConnection) error

	ListNodes(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConnection, error)
	ListConnections(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConnection, error)
	ListLeastUsedNodes(ctx context.Context, req domain.PickNodeReq) ([]domain.VelezConnection, error)
}

type Resources interface {
	Get(ctx context.Context, name string) (*domain.Resource, error)
	Create(ctx context.Context, resource domain.Resource) error

	Update(ctx context.Context, resource domain.Resource) error
}

type ConnectionCache interface {
	Add(nodes ...domain.Node)
	Get(names ...string) ([]domain.Node, error)
}
