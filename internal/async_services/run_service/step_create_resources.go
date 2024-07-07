package run_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

type CreateResourcesStep struct {
	resourceData storage.Resources
}

func NewCreateResourcesStep(resourceData storage.Resources) *CreateResourcesStep {
	return &CreateResourcesStep{
		resourceData: resourceData,
	}
}

func (c *CreateResourcesStep) Do(ctx context.Context, r *RunServiceReq) error {
	nextNode := loop_over.LoopOver(r.Nodes)

	for _, dependency := range r.Dependencies.Resources {
		resources, err := c.resourceData.Get(ctx, dependency.Name)
		if err != nil {
			return errors.Wrap(err, "error getting resource")
		}

		if len(resources) != 0 {
			continue
		}

		node := nextNode()

		var resource domain.Resource
		resource.ResourceName = dependency.Name
		resource.NodeName = node.Name

		err = c.resourceData.Create(ctx, resource)
		if err != nil {
			return errors.Wrap(err, "error creating resource")
		}

		resourceInstance, err := node.Conn.CreateSmerd(ctx, dependency.SmerdReq)
		if err != nil {
			return errors.Wrap(err, "error creating smerd on node")
		}

		if len(resourceInstance.Ports) == 0 {
			return errors.Wrap(ErrCreatedResourceHasNoPortsToAccess, "no ports returned")
		}

		resource.State = domain.ResourceStateCreated
		resource.Port = resourceInstance.Ports[0].Host

		err = c.resourceData.Update(ctx, resource)
		if err != nil {
			return errors.Wrap(err, "error changing state of resource")
		}
		listReq := &velez_api.ListSmerds_Request{
			Name: &resource.ResourceName,
		}
		resourcesList, err := node.Conn.ListSmerds(ctx, listReq)
		if err != nil {
			return errors.Wrap(err, "error checking for the state of resource")
		}

		var startedResource *velez_api.Smerd
		for _, runningResource := range resourcesList.Smerds {
			if runningResource.Name == resource.ResourceName {
				startedResource = runningResource
				break
			}
		}

		if startedResource == nil || startedResource.Status != velez_api.Smerd_running {
			resource.State = domain.ResourceStateError
		} else {
			resource.State = domain.ResourceStateRunning
		}

		err = c.resourceData.Update(ctx, resource)
		if err != nil {
			return errors.Wrap(err, "error updating resource state")
		}
	}
	return nil
}
