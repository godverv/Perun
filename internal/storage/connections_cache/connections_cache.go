package connections_cache

import (
	"context"
	"strconv"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Red-Sock/Perun/internal/domain"
)

const velezAuthHeader = "Velez-Auth"

type ConnectionConfig struct {
	Addr         string
	User         string
	SshKey       []byte
	VelezKeyPath string
}

type Velez struct {
	velez velez_api.VelezAPIClient
	md    metadata.MD
}

func NewVelezService(vConn domain.VelezConnection) (velez_api.VelezAPIClient, error) {
	addr := vConn.Node.Addr

	if vConn.Node.Port != nil {
		addr += ":" + strconv.Itoa(*vConn.Node.Port)
	}

	conn, err := getGrpcConnection(addr)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	srv := &Velez{
		velez: velez_api.NewVelezAPIClient(conn),
	}

	if vConn.Node.IsInsecure {
		return srv, nil
	}

	key, err := getVelezKey(vConn)
	if err != nil {
		return nil, errors.Wrap(err, "error getting velez key")
	}

	srv.md = metadata.New(
		map[string]string{
			velezAuthHeader: string(key),
		})

	return srv, nil
}

func (v *Velez) Version(ctx context.Context, in *velez_api.Version_Request, opts ...grpc.CallOption) (*velez_api.Version_Response, error) {
	ctx = metadata.NewOutgoingContext(ctx, v.md)

	return v.velez.Version(ctx, in, opts...)
}

func (v *Velez) CreateSmerd(ctx context.Context, in *velez_api.CreateSmerd_Request, opts ...grpc.CallOption) (*velez_api.Smerd, error) {
	ctx = metadata.NewOutgoingContext(ctx, v.md)
	return v.velez.CreateSmerd(ctx, in, opts...)
}

func (v *Velez) ListSmerds(ctx context.Context, in *velez_api.ListSmerds_Request, opts ...grpc.CallOption) (*velez_api.ListSmerds_Response, error) {
	ctx = metadata.NewOutgoingContext(ctx, v.md)
	return v.velez.ListSmerds(ctx, in, opts...)
}

func (v *Velez) DropSmerd(ctx context.Context, in *velez_api.DropSmerd_Request, opts ...grpc.CallOption) (*velez_api.DropSmerd_Response, error) {
	ctx = metadata.NewOutgoingContext(ctx, v.md)
	return v.velez.DropSmerd(ctx, in, opts...)
}

func (v *Velez) GetHardware(ctx context.Context, in *velez_api.GetHardware_Request, opts ...grpc.CallOption) (*velez_api.GetHardware_Response, error) {
	ctx = metadata.NewOutgoingContext(ctx, v.md)
	return v.velez.GetHardware(ctx, in, opts...)
}

func (v *Velez) FetchConfig(ctx context.Context, in *velez_api.FetchConfig_Request, opts ...grpc.CallOption) (*velez_api.FetchConfig_Response, error) {
	ctx = metadata.NewOutgoingContext(ctx, v.md)
	return v.velez.FetchConfig(ctx, in, opts...)
}
