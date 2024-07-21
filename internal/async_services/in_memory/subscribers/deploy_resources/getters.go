package deploy_resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (d *DeployResources) getResources(ctx context.Context, resourcesNames []string) (map[string]domain.Resource, error) {
	listResourcesReq := domain.DeployResourcesReq{
		ResourcesNames: resourcesNames,
	}
	resources, err := d.resourcesData.List(ctx, listResourcesReq)
	if err != nil {
		return nil, errors.Wrap(err, "error listing resources")
	}

	out := make(map[string]domain.Resource)

	for _, r := range resources {
		out[r.Name] = r
	}

	return out, nil
}

func (d *DeployResources) getDeployedResources(ctx context.Context, resourcesNames []string) (map[string]domain.Instance, error) {
	listInstancesReq := domain.ListInstancesReq{
		Names: resourcesNames,
	}
	deployedResources, err := d.instancesData.List(ctx, listInstancesReq)
	if err != nil {
		return nil, errors.Wrap(err, "error listing instances")
	}

	out := make(map[string]domain.Instance)
	for _, r := range deployedResources {
		out[r.Name] = r
	}

	return out, nil
}
