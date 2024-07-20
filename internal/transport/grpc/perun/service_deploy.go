package perun

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

func (s *Impl) DeployService(_ context.Context, req *perun_api.DeployService_Request) (
	*perun_api.DeployService_Response, error) {

	deployServiceReq := domain.DeployServiceReq{
		ServiceName: req.ServiceName,
	}

	err := s.deployServiceQueue.Dispatch(deployServiceReq)
	if err != nil {
		return nil, errors.Wrap(err, "error dispatching deploy service event")
	}

	return &perun_api.DeployService_Response{}, nil
}
