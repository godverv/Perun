package run_service_DEPRECATED

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

var ErrCreatedResourceHasNoPortsToAccess = errors.New("created resource has no ports to access")

type CreateResourcesStep struct {
	resourceData storage.Services
}

func NewCreateResourcesStep(resourceData storage.Services) *CreateResourcesStep {
	return &CreateResourcesStep{
		resourceData: resourceData,
	}
}

func (c *CreateResourcesStep) Do(ctx context.Context, r *RunServiceReq) error {
	nextNode := loop_over.LoopOver(r.Nodes)

	for _, dependency := range r.Dependencies.Resources {
		serviceDb, err := c.resourceData.Get(ctx, dependency.Name)
		if err != nil {
			return errors.Wrap(err, "error getting resource")
		}

		if serviceDb != nil {
			continue
		}

		node := nextNode()

		var resource domain.Service
		resource.Name = dependency.Name
		// TODO
		//resource.NodeName = node.Name

		err = c.resourceData.Upsert(ctx, resource)
		if err != nil {
			return errors.Wrap(err, "error creating resource")
		}

		resourceInstance, err := node.Conn.CreateSmerd(ctx, dependency.Constructor)
		if err != nil {
			return errors.Wrap(err, "error creating smerd on node")
		}

		if len(resourceInstance.Ports) == 0 {
			return errors.Wrap(ErrCreatedResourceHasNoPortsToAccess, "no ports returned")
		}

		resource.State = domain.ServiceStateCreated
		// TODO
		//resource.Port = resourceInstance.Ports[0].Host

		err = c.resourceData.UpdateState(ctx, resource)
		if err != nil {
			return errors.Wrap(err, "error changing state of resource")
		}
		listReq := &velez_api.ListSmerds_Request{
			Name: &resource.Name,
		}
		resourcesList, err := node.Conn.ListSmerds(ctx, listReq)
		if err != nil {
			return errors.Wrap(err, "error checking for the state of resource")
		}

		var startedResource *velez_api.Smerd
		for _, runningResource := range resourcesList.Smerds {
			if runningResource.Name == resource.Name {
				startedResource = runningResource
				break
			}
		}

		if startedResource == nil || startedResource.Status != velez_api.Smerd_running {
			resource.State = domain.ServiceStateError
		} else {
			resource.State = domain.ServiceStateRunning
		}

		err = c.resourceData.UpdateState(ctx, resource)
		if err != nil {
			return errors.Wrap(err, "error updating resource state")
		}
	}
	return nil
}
