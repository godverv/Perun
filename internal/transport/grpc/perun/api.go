package perun

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Red-Sock/Perun/internal/async_services"
	"github.com/Red-Sock/Perun/internal/config"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

type Impl struct {
	perun_api.UnimplementedPerunAPIServer

	nodeService service.NodesService

	initServiceQueue    async_services.ConsumerQueue[domain.InitServiceReq]
	refreshServiceQueue async_services.ConsumerQueue[domain.RefreshService]

	deployResourceQueue async_services.ConsumerQueue[domain.DeployResourcesReq]
	deployServiceQueue  async_services.ConsumerQueue[domain.DeployServiceReq]

	version string
}

func New(cfg config.Config, nodeService service.NodesService, queue async_services.AsyncService) *Impl {
	return &Impl{
		version:     cfg.GetAppInfo().Version,
		nodeService: nodeService,

		initServiceQueue:    queue.InitServiceQueue(),
		refreshServiceQueue: queue.RefreshServiceQueue(),

		deployResourceQueue: queue.DeployResourceQueue(),
		deployServiceQueue:  queue.DeployServiceQueue(),
	}
}

func (s *Impl) Register(server grpc.ServiceRegistrar) {
	perun_api.RegisterPerunAPIServer(server, s)
}

func (s *Impl) RegisterGw(ctx context.Context, mux *runtime.ServeMux, addr string) error {
	return perun_api.RegisterPerunAPIHandlerFromEndpoint(
		ctx,
		mux,
		addr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
}
