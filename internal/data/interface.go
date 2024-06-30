package data

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

type Storage interface {
	Nodes() Nodes
	Resources() Resources
}

type Nodes interface {
	GetConnection(ctx context.Context, key string) (*domain.VelezConnection, error)
	SaveConnection(ctx context.Context, in domain.VelezConnection) error

	ListNodes(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConn, error)
	ListConnections(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConnection, error)
}

type Resources interface {
	Get(ctx context.Context, name string) (*domain.Resource, error)
	Create(ctx context.Context, resource domain.Resource) error

	UpdateState(ctx context.Context, changeState domain.UpdateState) error
}
