package deploy_resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/domain"
)

type deployReq struct {
	resource domain.Resource
	node     domain.Node
}

func (d *DeployResources) deploy(ctx context.Context, deploy deployReq) error {
	resConstr, err := d.constructorData.GetByResourceName(ctx, deploy.resource.Name)
	if err != nil {
		return errors.Wrap(err, "error getting constructor data")
	}

	newInstance := domain.Instance{
		Name:     deploy.resource.Name,
		NodeName: deploy.node.Name,
		State:    domain.ServiceStateStarting,
		Image:    deploy.resource.Image,
	}
	// Stage 1 - pre-write
	{
		err = d.logDeployStarted(ctx, newInstance)
		if err != nil {
			return errors.Wrap(err)
		}

		err = d.instancesData.Add(ctx, newInstance)
		if err != nil {
			return errors.Wrap(err, "error saving instance info")
		}
	}

	// Stage 2 - deploy
	{
		var workingNodeResponse *velez_api.Smerd
		workingNodeResponse, err = deploy.node.Conn.CreateSmerd(ctx, resConstr.DeployRequest)
		if err != nil {
			return errors.Wrap(err, "error creating smerd on node")
		}

		if len(workingNodeResponse.Ports) == 0 {
			return d.logWorkingNodeReturnedZeroPorts(ctx, newInstance)
		}
		newInstance.Port = int(workingNodeResponse.Ports[0].Host)
		newInstance.State = domain.ServiceStateRunningOk
	}

	// Stage 3 - commit deploy
	{
		err = d.instancesData.Update(ctx, newInstance)
		if err != nil {
			return errors.Wrap(err, "error saving instance info")
		}

		err = d.logDeploySuccessful(ctx, newInstance)
		if err != nil {
			return errors.Wrap(err)
		}

	}

	return nil
}
