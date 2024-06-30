package nodes_service

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (n *nodesService) ListNodes(ctx context.Context, nodes domain.ListVelezNodes) ([]domain.VelezConn, error) {
	return n.nodesStore.ListNodes(ctx, nodes)
}
