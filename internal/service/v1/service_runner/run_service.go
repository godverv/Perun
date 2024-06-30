package service_runner

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"
	"golang.org/x/sync/errgroup"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

func (r *serviceRunner) Run(ctx context.Context, req domain.RunServiceRequest) error {
	if len(req.Nodes) == 0 {
		return errors.New("no nodes to run service on")
	}

	err := r.syncConfig(ctx, req)
	if err != nil {
		return errors.Wrap(err, "error syncing config")
	}

	err = r.syncDependencies(ctx, req)
	if err != nil {
		return errors.Wrap(err, "error syncing dependencies")
	}

	g, ctx := errgroup.WithContext(ctx)

	for _, node := range req.Nodes {
		g.Go(func() error {
			_, err := node.Conn.CreateSmerd(ctx, req.Constructor)
			return err
		})
	}

	err = g.Wait()
	if err != nil {
		return errors.Wrap(err, "error creating smerds on nodes")
	}

	return nil
}

func (r *serviceRunner) runService(
	ctx context.Context,
	cl velez_api.VelezAPIClient,
	req *velez_api.CreateSmerd_Request,
) error {
	_, err := cl.CreateSmerd(ctx, req)
	if err != nil {
		return errors.Wrap(err, "error creating smerd")
	}
	return nil
}

func (r *serviceRunner) syncConfig(ctx context.Context, req domain.RunServiceRequest) error {
	fetchReq := &velez_api.FetchConfig_Request{
		ImageName:   req.Constructor.ImageName,
		ServiceName: req.Constructor.Name,
	}
	_, err := req.Nodes[0].Conn.FetchConfig(ctx, fetchReq)
	if err != nil {
		return errors.Wrap(err, "error fetching config")
	}

	return nil
}

func (r *serviceRunner) syncDependencies(ctx context.Context, req domain.RunServiceRequest) error {
	dependencies, err := r.resourceService.GetDependencies(ctx, req.Constructor.Name)
	if err != nil {
		return errors.Wrap(err, "error getting dependencies")
	}

	if req.Constructor.Settings == nil {
		req.Constructor.Settings = &velez_api.Container_Settings{}
	}
	req.Constructor.Settings.Volumes = append(req.Constructor.Settings.Volumes, dependencies.Volumes...)

	nextNode := loop_over.LoopOver(req.Nodes)

	for _, dep := range dependencies.Smerds {
		res, err := r.resourceData.Get(ctx, dep.Name)
		if err != nil {
			return errors.Wrap(err, "error getting dependency from db")
		}

		if res == nil {
			res, err = r.createResource(ctx, nextNode(), dep)
			if err != nil {
				return errors.Wrap(err, "error creating resource")
			}
		}
	}

	return nil
}

func (r *serviceRunner) createResource(ctx context.Context, node domain.Node, req *velez_api.CreateSmerd_Request) (
	*domain.Resource, error) {

	resource := domain.Resource{
		ResourceName: req.Name,
		NodeName:     node.Name,
	}
	err := r.resourceData.Create(ctx, resource)
	if err != nil {
		return nil, errors.Wrap(err, "error creating resource")
	}

	_, err = node.Conn.CreateSmerd(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error creating smerd on node")
	}

	updateStateReq := domain.UpdateState{
		ResourceName: req.Name,
		State:        domain.ResourceStateCreated,
	}
	err = r.resourceData.UpdateState(ctx, updateStateReq)
	if err != nil {
		return nil, errors.Wrap(err, "error changing state of resource")
	}

	return &resource, nil
}
