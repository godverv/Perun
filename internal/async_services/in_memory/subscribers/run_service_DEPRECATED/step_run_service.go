package run_service_DEPRECATED

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

type RunServiceStep struct {
	resourceData storage.Services
}

func NewRunServiceStep(resourceData storage.Services) *RunServiceStep {
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

		// TODO
		createResource := domain.Service{
			Name: req.ServiceName,
			//NodeName: node.Name,
		}
		err := r.resourceData.Upsert(ctx, createResource)
		if err != nil {
			return errors.Wrap(err, "error creating resource")
		}

		createSmerd := &velez_api.CreateSmerd_Request{
			Name:      req.ServiceName,
			ImageName: req.ImageName,
		}

		resourceInstance, err := node.Conn.CreateSmerd(ctx, createSmerd)
		if err != nil {
			return errors.Wrap(err, "error creating smerd on node")
		}

		updateStateReq := domain.Service{
			Name: req.ServiceName,
			//NodeName: node.Name,
			State: domain.ServiceStateCreated,
		}
		if len(resourceInstance.Ports) != 0 {
			// TODO
			//updateStateReq.Port = resourceInstance.Ports[0].Host
		}

		err = r.resourceData.UpdateState(ctx, updateStateReq)
		if err != nil {
			return errors.Wrap(err, "error changing state of resource")
		}
	}

	return nil
}
