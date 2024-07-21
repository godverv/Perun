package deploy_resources

import (
	"context"
	stderrors "errors"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type deployReq struct {
	resource domain.Resource
	node     domain.Node

	deployLogSrv   service.DeployLog
	instancesData  storage.Instances
	deployPatterns storage.DeployPatterns
}

func (d *deployReq) do(ctx context.Context) (err error) {
	var inst domain.Instance

	defer func() {
		if err == nil {
			return
		}

		fallbackErr := d.fallback(ctx, inst)
		if fallbackErr != nil {
			err = stderrors.Join(err, errors.Wrap(fallbackErr))
		}
	}()

	inst, err = d.prepare(ctx)
	if err != nil {
		err = errors.Wrap(err)
		return
	}

	inst, err = d.deployToNode(ctx, inst)
	if err != nil {
		return errors.Wrap(err)
	}

	err = d.commitDeployment(ctx, inst)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Stage 1 - pre-write
func (d *deployReq) prepare(ctx context.Context) (domain.Instance, error) {
	newInstance := domain.Instance{
		Name:     d.resource.Name,
		NodeName: d.node.Name,
		State:    domain.ServiceStateStarting,
		Image:    d.resource.Image,
	}

	err := d.deployLogSrv.DeployStarted(ctx, newInstance)
	if err != nil {
		return newInstance, errors.Wrap(err)
	}
	err = d.instancesData.Add(ctx, newInstance)
	if err != nil {
		return newInstance, errors.Wrap(err, "error saving instance info")
	}

	return newInstance, nil
}

// Stage 2 - deploy to node
func (d *deployReq) deployToNode(ctx context.Context, in domain.Instance) (domain.Instance, error) {
	resConstr, err := d.deployPatterns.GetByResourceName(ctx, d.resource.Name)
	if err != nil {
		return in, errors.Wrap(err, "error getting constructor data")
	}

	workingNodeResponse, err := d.node.Conn.CreateSmerd(ctx, resConstr.DeployRequest)
	if err != nil {
		return in, errors.Wrap(err, "error creating smerd on node")
	}

	if len(workingNodeResponse.Ports) == 0 {
		return in, d.deployLogSrv.WorkingNodeReturnedZeroPorts(ctx, in)
	}

	in.Port = int(workingNodeResponse.Ports[0].Host)
	in.State = domain.ServiceStateRunningOk
	return in, nil
}

// Stage 3 - commit deploy
func (d *deployReq) commitDeployment(ctx context.Context, inst domain.Instance) error {
	err := d.instancesData.Update(ctx, inst)
	if err != nil {
		return errors.Wrap(err, "error saving instance info")
	}

	err = d.deployLogSrv.DeploySuccessful(ctx, inst)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (d *deployReq) fallback(ctx context.Context, inst domain.Instance) error {
	dropSmerdReq := &velez_api.DropSmerd_Request{
		Name: []string{inst.Name},
	}
	var err error
	_, dropSmerdErr := d.node.Conn.DropSmerd(ctx, dropSmerdReq)
	if dropSmerdErr != nil {
		dropSmerdErr = errors.Wrap(dropSmerdErr, "error dropping smerd instance")
		err = stderrors.Join(err, dropSmerdErr)
	}

	deleteInstanceErr := d.instancesData.Delete(ctx, inst.Name)
	if deleteInstanceErr != nil {
		deleteInstanceErr = errors.Wrap(deleteInstanceErr, "error deleting failed instance")
		err = stderrors.Join(err, deleteInstanceErr)
	}

	logErr := d.deployLogSrv.DeleteFailedDeployment(ctx, inst)
	if logErr != nil {
		logErr = errors.Wrap(logErr, "error deleting failed deployment")
		err = stderrors.Join(err, logErr)
	}

	return err
}
