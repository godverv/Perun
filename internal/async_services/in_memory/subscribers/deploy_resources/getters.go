package deploy_resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/utils/loop_over"
)

func (d *DeployResources) getResources(ctx context.Context, resourcesNames []string) (map[string]domain.Resource, error) {
	listResourcesReq := domain.DeployResourcesReq{
		ResourcesNames: resourcesNames,
	}
	resources, err := d.resourcesData.List(ctx, listResourcesReq)
	if err != nil {
		return nil, errors.Wrap(err, "error listing resources")
	}

	if len(resources) == 0 {
		return nil, errors.New("resources not found")
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

func (d *DeployResources) getNodeIterator(ctx context.Context, nodesCount uint32) (func() domain.Node, error) {
	pickNodes := domain.PickNodesReq{
		NodesCount: nodesCount,
	}
	n, err := d.nodes.PickNodes(ctx, pickNodes)
	if err != nil {
		return nil, errors.Wrap(err, "error picking nodes")
	}

	nextNode := loop_over.LoopOver(n)
	return nextNode, nil
}
