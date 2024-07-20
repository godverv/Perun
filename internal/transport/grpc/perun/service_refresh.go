package perun

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	api "github.com/Red-Sock/Perun/pkg/perun_api"
)

func (s *Impl) RefreshService(ctx context.Context, req *api.RefreshService_Request,
) (*api.RefreshService_Response, error) {

	refreshReq := domain.RefreshService{
		ServiceName: req.ServiceName,
	}

	err := s.refreshServiceQueue.Dispatch(refreshReq)
	if err != nil {
		return nil, errors.Wrap(err, "error putting event into queue")
	}

	return &api.RefreshService_Response{}, nil
}
