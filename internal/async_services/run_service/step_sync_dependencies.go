package run_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

var ErrCreatedResourceHasNoPortsToAccess = errors.New("created resource has no ports to access")

type SyncDependenciesStep struct {
	resourceService service.ResourceService
	resourceData    data.Resources
}

func NewSyncDependenciesStep(
	resourceService service.ResourceService,
	resourceData data.Resources,
) *SyncDependenciesStep {
	return &SyncDependenciesStep{
		resourceService: resourceService,
		resourceData:    resourceData,
	}
}

func (s *SyncDependenciesStep) Do(ctx context.Context, r *RunServiceReq) error {
	if len(r.Nodes) == 0 {
		return errors.Wrap(ErrNoNodes, "no nodes provided to sync dependencies")
	}

	dependencies, err := s.resourceService.GetDependencies(ctx, r.ServiceName)
	if err != nil {
		return errors.Wrap(err, "error getting dependencies")
	}

	nextNode := loop_over.LoopOver(r.Nodes)

	for _, dep := range dependencies.Smerds {
		var res *domain.Resource
		res, err = s.resourceData.Get(ctx, dep.Name)
		if err != nil {
			return errors.Wrap(err, "error getting dependency from db")
		}

		if res == nil {
			// todo: when creating resource 100% there is gonna be change in config
			err = s.createResource(ctx, nextNode(), dep)
			if err != nil {
				return errors.Wrap(err, "error creating resource")
			}
		}

		// todo: handle current resource state
	}

	return nil
}

func (s *SyncDependenciesStep) createResource(ctx context.Context, node domain.Node, req domain.Dependency) error {
	var createResReq domain.Resource
	createResReq.ResourceName = req.Name
	createResReq.NodeName = node.Name

	err := s.resourceData.Create(ctx, createResReq)
	if err != nil {
		return errors.Wrap(err, "error creating resource")
	}

	resourceInstance, err := node.Conn.CreateSmerd(ctx, req.SmerdReq)
	if err != nil {
		return errors.Wrap(err, "error creating smerd on node")
	}

	if len(resourceInstance.Ports) == 0 {
		return errors.Wrap(ErrCreatedResourceHasNoPortsToAccess, "no ports returned")
	}
	updateStateReq := domain.Resource{
		ResourceName: req.Name,
		NodeName:     node.Name,
		State:        domain.ResourceStateCreated,
		Port:         resourceInstance.Ports[0].Host,
	}

	err = s.resourceData.Update(ctx, updateStateReq)
	if err != nil {
		return errors.Wrap(err, "error changing state of resource")
	}

	return nil
}
