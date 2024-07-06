package run_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
)

var (
	ErrNoNodes = errors.New("no nodes")
)

type GetNodesStep struct {
	nodes service.NodesService
}

func NewGetNodesStep(nodes service.NodesService) *GetNodesStep {
	return &GetNodesStep{
		nodes: nodes,
	}
}

func (g *GetNodesStep) Do(ctx context.Context, r *RunServiceReq) (err error) {
	var pickNodesReq domain.PickNodeReq
	pickNodesReq.ReplicationFactor = r.ReplicationFactor

	r.Nodes, err = g.nodes.PickNodes(ctx, pickNodesReq)
	if err != nil {
		return errors.Wrap(err, "error getting nodes")
	}

	if len(r.Nodes) == 0 {
		return errors.Wrap(ErrNoNodes, "not found")
	}

	return nil
}
