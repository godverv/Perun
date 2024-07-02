package nodes_service

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (n *NodesService) ListNodes(ctx context.Context, nodes domain.ListVelezNodes) ([]domain.VelezConnection, error) {
	return n.nodesStore.ListNodes(ctx, nodes)
}
