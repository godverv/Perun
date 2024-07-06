package run_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

type RunServiceStep struct {
	resourceData data.Resources
}

func NewRunServiceStep(resourceData data.Resources) *RunServiceStep {
	return &RunServiceStep{
		resourceData: resourceData,
	}
}

func (r *RunServiceStep) Do(ctx context.Context, req *RunServiceReq) error {
	if uint32(len(req.Nodes)) < req.ReplicationFactor {
		return errors.Wrap(ErrNoNodes, "no nodes to run service on")
	}

	nextNode := loop_over.LoopOver(req.Nodes)

	for i := uint32(0); i < req.ReplicationFactor; i++ {
		node := nextNode()

		createResource := domain.Resource{
			ResourceName: req.ServiceName,
			NodeName:     node.Name,
		}
		err := r.resourceData.Create(ctx, createResource)
		if err != nil {
			return errors.Wrap(err, "error creating resource")
		}

		createSmerd := &velez_api.CreateSmerd_Request{
			Name:      req.ServiceName,
			ImageName: req.ImageName,
			Env:       make(map[string]string), // TODO: pass credentials for resources
		}

		resourceInstance, err := node.Conn.CreateSmerd(ctx, createSmerd)
		if err != nil {
			return errors.Wrap(err, "error creating smerd on node")
		}

		if len(resourceInstance.Ports) == 0 {
			return errors.Wrap(ErrCreatedResourceHasNoPortsToAccess, "services didn't return any ports for deployed service")
		}

		updateStateReq := domain.Resource{
			ResourceName: req.ServiceName, // TODO handle multiple instances on different nodes
			NodeName:     node.Name,
			State:        domain.ResourceStateCreated,
			Port:         resourceInstance.Ports[0].Host,
		}

		err = r.resourceData.Update(ctx, updateStateReq)
		if err != nil {
			return errors.Wrap(err, "error changing state of resource")
		}
	}

	return nil
}
