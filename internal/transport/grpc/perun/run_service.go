package perun

import (
	"context"

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
	impl.createServiceQ <- runServiceReq

	return &perun_api.RunService_Response{}, nil
}
