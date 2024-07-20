package refresh_service_config

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type RefreshServiceConfig struct {
	nodesService service.NodesService

	servicesData  storage.Services
	resourcesData storage.Resources
}

func New(data storage.Data, srv service.Services) *RefreshServiceConfig {
	return &RefreshServiceConfig{
		nodesService: srv.Nodes(),

		servicesData:  data.Services(),
		resourcesData: data.Resources(),
	}
}

func (r *RefreshServiceConfig) Consume(ctx context.Context, req domain.RefreshService) error {
	srv, err := r.servicesData.Get(ctx, req.ServiceName)
	if err != nil {
		return errors.Wrap(err, "error getting service info from storage")
	}

	cfg, err := r.syncConfig(ctx, srv)
	if err != nil {
		return errors.Wrap(err, "error syncing config")
	}

	err = r.syncDependencies(ctx, req, cfg)
	if err != nil {
		return errors.Wrap(err, "error syncing dependencies")
	}

	return nil
}
