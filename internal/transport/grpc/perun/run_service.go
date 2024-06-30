package perun

import (
	"context"

	"github.com/godverv/Velez/pkg/velez_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

func (impl *Implementation) RunService(
	ctx context.Context, req *perun_api.RunService_Request) (*perun_api.RunService_Response, error) {

	nodes, err := impl.pickNodes(ctx, req)
	if err != nil {
		return nil, err
	}

	runServiceReq := domain.RunServiceRequest{
		Constructor: &velez_api.CreateSmerd_Request{
			Name:      req.ServiceName,
			ImageName: req.ImageName,
		},
		Nodes: nodes,
	}
	err = impl.runService.Run(ctx, runServiceReq)
	if err != nil {
		return nil, err
	}

	return &perun_api.RunService_Response{}, nil
}

func (impl *Implementation) pickNodes(ctx context.Context, req *perun_api.RunService_Request) ([]domain.Node, error) {
	pickNodeReq := domain.PickNodeReq{
		ServiceName:       req.ServiceName,
		ReplicationFactor: req.ReplicationFactor,
	}
	nodes, err := impl.nodeService.PickNodes(ctx, pickNodeReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if len(nodes) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "no availabel nodes")
	}

	return nodes, nil
}
