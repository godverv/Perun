package deploy_log

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *DeployLogService) WorkingNodeReturnedZeroPorts(ctx context.Context, res domain.Instance) error {
	log := domain.DeployLog{
		Name:   res.Name,
		State:  domain.ServiceStateErrorDuringDeploy,
		Reason: "Resource " + res.Name + " was deployed, but no ports have returned from node " + res.NodeName,
	}

	err := s.deployLogData.Add(ctx, log)
	if err != nil {
		return errors.Wrap(err, "error logging deploy error")
	}

	return nil
}
