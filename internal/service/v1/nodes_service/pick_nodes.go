package nodes_service

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (n *nodesService) PickNodes(_ context.Context, req domain.PickNodeReq) ([]domain.Node, error) {
	nodes := make([]domain.Node, 0, req.ReplicationFactor)

	i := uint32(0)
	for _, node := range n.nodes {
		if i == req.ReplicationFactor {
			break
		}

		nodes = append(nodes, node)
		i++
	}

	return nodes, nil
}
