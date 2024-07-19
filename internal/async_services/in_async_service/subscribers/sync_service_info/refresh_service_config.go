package sync_service_info

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/async_services"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type RefreshServiceConfig struct {
	nodesService service.NodesService

	servicesData storage.Services

	syncEventBus async_services.Queue[domain.SyncServiceInfo]
}

func New(data storage.Data, srv service.Services) *RefreshServiceConfig {
	return &RefreshServiceConfig{
		nodesService: srv.Nodes(),

		servicesData: data.Services(),
	}
}

func (r *RefreshServiceConfig) Consume(ctx context.Context, req domain.SyncServiceInfo) error {
	cfg, err := r.syncConfig(ctx, req)
	if err != nil {
		return errors.Wrap(err, "error syncing config")
	}

	err = r.syncDependencies(ctx, req, cfg)
	if err != nil {
		return errors.Wrap(err, "error syncing dependencies")
	}

	// TODO create dep in db
	return nil
}
