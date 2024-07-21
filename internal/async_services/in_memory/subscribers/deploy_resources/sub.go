package deploy_resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"golang.org/x/sync/errgroup"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type DeployResources struct {
	nodes     service.NodesService
	config    service.ConfigService
	deployLog service.DeployLog

	resourcesData  storage.Resources
	instancesData  storage.Instances
	serviceData    storage.Services
	deployPatterns storage.DeployPatterns
}

func New(data storage.Data, srv service.Services) *DeployResources {
	return &DeployResources{
		resourcesData:  data.Resources(),
		instancesData:  data.Instances(),
		serviceData:    data.Services(),
		deployPatterns: data.DeployTemplates(),

		deployLog: srv.DeployLog(),
		nodes:     srv.Nodes(),
		config:    srv.Config(),
	}
}

func (d *DeployResources) Consume(ctx context.Context, req domain.DeployResourcesReq) error {
	registeredResourcesMap, err := d.getResources(ctx, req.ResourcesNames)
	if err != nil {
		return errors.Wrap(err)
	}

	deployedResourcesMap, err := d.getDeployedResources(ctx, req.ResourcesNames)
	if err != nil {
		return errors.Wrap(err)
	}

	nextNode, err := d.getNodeIterator(ctx, uint32(len(registeredResourcesMap)-len(deployedResourcesMap)))
	if err != nil {
		return errors.Wrap(err)
	}

	g, gctx := errgroup.WithContext(ctx)
	for _, registeredRes := range registeredResourcesMap {

		deployed, ok := deployedResourcesMap[registeredRes.Name]
		if ok && deployed.State == domain.ServiceStateRunningOk {
			g.Go(func() error {
				return d.deployLog.AlreadyDeployed(gctx, deployed)
			})
		} else {
			dp := d.newDeploy(registeredRes, nextNode())
			g.Go(func() error {
				return dp.do(ctx)
			})
		}
	}

	err = g.Wait()
	if err != nil {
		return errors.Wrap(err, "error during deploy")
	}

	return nil
}

func (d *DeployResources) newDeploy(res domain.Resource, node domain.Node) deployReq {
	return deployReq{
		resource: res,
		node:     node,

		deployLogSrv:   d.deployLog,
		instancesData:  d.instancesData,
		deployPatterns: d.deployPatterns,
	}
}
