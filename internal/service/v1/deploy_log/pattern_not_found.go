package deploy_log

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *DeployLogService) DeployPatternNotFounds(ctx context.Context, res domain.Resource) error {
	log := domain.DeployLog{
		Name:   res.Name,
		State:  domain.ServiceStateErrorDuringDeploy,
		Reason: "No constructor found for image: " + res.Image,
	}

	err := s.deployLogData.Add(ctx, log)
	if err != nil {
		return errors.Wrap(err, "error logging deploy error")
	}

	return nil
}
