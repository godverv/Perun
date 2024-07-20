package nodes_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

var ErrNoNodes = errors.New("no nodes")

func (n *NodesService) PickNodes(ctx context.Context, req domain.PickNodesReq) ([]domain.Node, error) {
	velezConnections, err := n.nodesStore.ListLeastUsedNodes(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error listing velezConnections")
	}

	nodesNames := make([]string, 0, len(velezConnections))
	for _, node := range velezConnections {
		nodesNames = append(nodesNames, node.Node.Name)
	}

	nodes, err := n.connectionsCache.Get(nodesNames...)
	if err != nil {
		return nil, errors.Wrap(err, "error getting nodes by name from cache")
	}

	if len(nodes) == 0 {
		return nil, errors.Wrap(ErrNoNodes)
	}

	return nodes, nil
}
