package perun

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

var ErrEmptyRequest = errors.New("empty request")

func (impl *Implementation) ConnectVelez(ctx context.Context, req *perun_api.ConnectVelez_Request) (*perun_api.ConnectVelez_Response, error) {
	velezNode, err := fromVelezNode(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = impl.nodeService.NewVelezConnection(ctx, velezNode)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &perun_api.ConnectVelez_Response{}, nil
}

func fromVelezNode(in *perun_api.ConnectVelez_Request) (domain.VelezConnection, error) {
	if in == nil {
		return domain.VelezConnection{}, ErrEmptyRequest
	}

	return domain.VelezConnection{
		Node: domain.Velez{
			Name:               in.GetNode().GetName(),
			Addr:               in.GetNode().GetAddr(),
			CustomVelezKeyPath: in.GetNode().GetCustomVelezKeyPath(),
			IsInsecure:         in.GetNode().GetSecurityDisabled(),
		},
		Ssh: domain.Ssh{
			Key:      in.GetSsh().GetKeyBase64(),
			Port:     in.GetSsh().GetPort(),
			Username: in.GetSsh().GetUsername(),
		},
	}, nil
}
