package perun

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

var (
	ErrEmptyRequest = errors.New("empty request")
)

func (s *Impl) ConnectVelez(ctx context.Context, req *perun_api.ConnectVelez_Request) (*perun_api.ConnectVelez_Response, error) {
	velezNode, err := fromVelezConn(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.nodeService.NewVelezConnection(ctx, velezNode)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &perun_api.ConnectVelez_Response{}, nil
}

func fromVelezConn(in *perun_api.ConnectVelez_Request) (vConn domain.VelezConnection, err error) {
	if in == nil {
		return domain.VelezConnection{}, ErrEmptyRequest
	}

	vConn.Node = fromNode(in.GetNode())
	vConn.Ssh = fromSsh(in.GetSsh())

	return vConn, nil
}

func fromNode(in *perun_api.Node) domain.Velez {
	var vConn domain.Velez

	vConn = domain.Velez{
		Name: in.GetName(),
		Addr: in.GetAddr(),

		CustomVelezKeyPath: in.GetCustomVelezKeyPath(),
		IsInsecure:         in.GetSecurityDisabled(),
	}

	if in.Port != nil {
		vPortUint := int(*in.Port)
		vConn.Port = &vPortUint
	}

	return vConn
}

func fromSsh(in *perun_api.Ssh) *domain.Ssh {
	if in == nil {
		return nil
	}
	return &domain.Ssh{
		Key:      in.GetKeyBase64(),
		Port:     in.GetPort(),
		Username: in.GetUsername(),
	}
}
