package perun

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Red-Sock/Perun/internal/async_services/run_service"
	"github.com/Red-Sock/Perun/internal/config"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

type Implementation struct {
	perun_api.UnimplementedPerunAPIServer

	nodeService     service.NodesService
	resourceService service.ResourceService

	createServiceQ chan<- run_service.RunServiceReq

	version string
}

func New(cfg config.Config, nodeService service.NodesService, createServiceQ chan<- run_service.RunServiceReq) *Implementation {
	return &Implementation{
		version:        cfg.GetAppInfo().Version,
		nodeService:    nodeService,
		createServiceQ: createServiceQ,
	}
}

func (impl *Implementation) Register(server grpc.ServiceRegistrar) {
	perun_api.RegisterPerunAPIServer(server, impl)
}

func (impl *Implementation) RegisterGw(ctx context.Context, mux *runtime.ServeMux, addr string) error {
	return perun_api.RegisterPerunAPIHandlerFromEndpoint(
		ctx,
		mux,
		addr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
}
