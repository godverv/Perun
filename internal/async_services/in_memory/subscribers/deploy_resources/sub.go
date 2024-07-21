package deploy_resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"golang.org/x/sync/errgroup"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

type DeployResources struct {
	nodes  service.NodesService
	config service.ConfigService

	resourcesData   storage.Resources
	instancesData   storage.Instances
	deployLog       storage.DeployLogs
	serviceData     storage.Services
	constructorData storage.ResourceConstructors
}

func New(data storage.Data, srv service.Services) *DeployResources {
	return &DeployResources{
		resourcesData:   data.Resources(),
		instancesData:   data.Instances(),
		deployLog:       data.DeployLogs(),
		serviceData:     data.Services(),
		constructorData: data.DeployTemplates(),

		nodes:  srv.Nodes(),
		config: srv.Config(),
	}
}

func (d *DeployResources) Consume(ctx context.Context, req domain.DeployResourcesReq) error {
	registeredResourcesMap, err := d.getResources(ctx, req.ResourcesNames)
	if err != nil {
		return errors.Wrap(err)
	}

	if len(registeredResourcesMap) == 0 {
		// TODO
		return nil
	}

	deployedResourcesMap, err := d.getDeployedResources(ctx, req.ResourcesNames)
	if err != nil {
		return errors.Wrap(err)
	}

	pickNodes := domain.PickNodesReq{
		NodesCount: uint32(len(registeredResourcesMap)),
	}
	n, err := d.nodes.PickNodes(ctx, pickNodes)
	if err != nil {
		return errors.Wrap(err, "error picking nodes")
	}

	nextNode := loop_over.LoopOver(n)

	g, gctx := errgroup.WithContext(ctx)
	for _, registeredRes := range registeredResourcesMap {
		deployed, ok := deployedResourcesMap[registeredRes.Name]
		if ok {
			// TODO Service might be registered as instance but not running
			// in this case (if failed) should clean mess and redeploy
			g.Go(func() error {
				return d.logAlreadyDeployed(gctx, deployed)
			})
		} else {
			dr := deployReq{
				resource: registeredRes,
				node:     nextNode(),
			}
			g.Go(func() error {
				return d.deploy(gctx, dr)
			})
		}
	}

	err = g.Wait()
	if err != nil {
		return errors.Wrap(err, "error during deploy")
	}

	return nil
}
