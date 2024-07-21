package deploy_log

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *DeployLogService) DeployStarted(ctx context.Context, inst domain.Instance) error {
	createNewInstanceLog := domain.DeployLog{
		Name:   inst.Name,
		State:  domain.ServiceStateStarting,
		Reason: "Deploying " + inst.Name + " on " + inst.NodeName,
	}

	err := s.deployLogData.Add(ctx, createNewInstanceLog)
	if err != nil {
		return errors.Wrap(err, "error logging start")
	}

	return nil
}
