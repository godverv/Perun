package deploy_log

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *DeployLogService) AlreadyDeployed(ctx context.Context, inst domain.Instance) error {
	log := domain.DeployLog{
		Name:   inst.Name,
		State:  domain.ServiceStateErrorAlreadyDeployed,
		Reason: "Deploy stopped: resources is already running on " + inst.NodeName,
	}

	err := s.deployLogData.Add(ctx, log)
	if err != nil {
		return errors.Wrap(err, "error logging deploy error")
	}

	return nil
}
