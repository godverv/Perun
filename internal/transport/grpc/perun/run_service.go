package perun

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Red-Sock/Perun/internal/async_services/run_service"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

func (impl *Implementation) RunService(
	ctx context.Context, req *perun_api.RunService_Request) (
	*perun_api.RunService_Response, error) {

	runServiceReq := run_service.RunServiceReq{
		ServiceName:       req.ServiceName,
		ImageName:         req.ImageName,
		ReplicationFactor: req.ReplicationFactor,
	}
	err := impl.runService.Run(ctx, runServiceReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &perun_api.RunService_Response{}, nil
}
