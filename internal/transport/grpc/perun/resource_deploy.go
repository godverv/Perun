package perun

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

func (s *Impl) DeployResource(_ context.Context, req *perun_api.DeployResource_Request) (
	*perun_api.DeployResource_Response, error) {

	deployReq := domain.DeployResourcesReq{
		ResourcesNames: req.GetResourceNames(),
	}
	err := s.deployResourceQueue.Dispatch(deployReq)
	if err != nil {
		return nil, errors.Wrap(err, "error dispatching event to queue")
	}

	return &perun_api.DeployResource_Response{}, nil
}
