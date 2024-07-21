package deploy_resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (d *DeployResources) logAlreadyDeployed(ctx context.Context, resource domain.Instance) error {
	log := domain.DeployLog{
		Name:   resource.Name,
		State:  domain.ErrorAlreadyDeployed,
		Reason: "Deploy stopped: resources is already running on " + resource.NodeName,
	}

	err := d.deployLog.Add(ctx, log)
	if err != nil {
		return errors.Wrap(err, "error logging deploy error")
	}

	return nil
}

func (d *DeployResources) logConstructorNotFounds(ctx context.Context, res domain.Resource) error {
	log := domain.DeployLog{
		Name:   res.Name,
		State:  domain.ErrorDuringDeploy,
		Reason: "No constructor found for image: " + res.Image,
	}

	err := d.deployLog.Add(ctx, log)
	if err != nil {
		return errors.Wrap(err, "error logging deploy error")
	}

	return nil
}

func (d *DeployResources) logWorkingNodeReturnedZeroPorts(ctx context.Context, res domain.Instance) error {
	log := domain.DeployLog{
		Name:   res.Name,
		State:  domain.ErrorDuringDeploy,
		Reason: "Resource " + res.Name + " was deployed, but no ports have returned from node " + res.NodeName,
	}

	err := d.deployLog.Add(ctx, log)
	if err != nil {
		return errors.Wrap(err, "error logging deploy error")
	}

	return nil
}

func (d *DeployResources) logDeployStarted(ctx context.Context, inst domain.Instance) error {
	createNewInstanceLog := domain.DeployLog{
		Name:   inst.Name,
		State:  domain.ServiceStateStarting,
		Reason: "Deploying " + inst.Name + " on " + inst.NodeName,
	}
	err := d.deployLog.Add(ctx, createNewInstanceLog)
	if err != nil {
		return errors.Wrap(err, "error logging start")
	}

	return nil
}

func (d *DeployResources) logDeploySuccessful(ctx context.Context, inst domain.Instance) error {
	createNewInstanceLog := domain.DeployLog{
		Name:   inst.Name,
		State:  domain.ServiceStateRunningOk,
		Reason: "Resource " + inst.Name + " successfully deployed on " + inst.NodeName,
	}
	err := d.deployLog.Add(ctx, createNewInstanceLog)
	if err != nil {
		return errors.Wrap(err, "error logging start")
	}

	return nil
}
