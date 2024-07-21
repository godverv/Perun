package deploy_log

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *DeployLogService) DeleteFailedDeployment(ctx context.Context, inst domain.Instance) error {
	deployLog := domain.DeployLog{
		Name:  inst.Name,
		State: domain.ServiceStateDeployDeleted,
	}

	err := s.deployLogData.Add(ctx, deployLog)
	if err != nil {
		return errors.Wrap(err, "error adding deploy log to db")
	}

	return nil
}
