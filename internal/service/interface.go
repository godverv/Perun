package service

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

type Services interface {
	Nodes() NodesService
}

type NodesService interface {
	NewVelezConnection(ctx context.Context, node domain.VelezConnection) error

	ListNodes(ctx context.Context, nodes domain.ListVelezNodes) ([]domain.VelezConnection, error)
	PickNodes(ctx context.Context, req domain.PickNodesReq) ([]domain.Node, error)
}
