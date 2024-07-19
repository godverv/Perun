package perun

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	api "github.com/Red-Sock/Perun/pkg/perun_api"
)

func (s *Impl) CreateService(_ context.Context, req *api.CreateService_Request) (
	*api.CreateService_Response, error) {

	initServiceReq := domain.InitServiceReq{
		ServiceName:       req.ServiceName,
		ImageName:         req.ImageName,
		ReplicationFactor: int(req.Replicas),
	}
	err := s.initServiceQueue.Dispatch(initServiceReq)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.CreateService_Response{}, nil
}
