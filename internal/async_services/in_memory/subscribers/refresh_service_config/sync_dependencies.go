package refresh_service_config

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service/v1/constructors/patterns"
)

func (r *RefreshServiceConfig) syncDependencies(ctx context.Context,
	req domain.RefreshService,
	config matreshka.AppConfig,
) error {
	depsList, err := r.resourcesData.ListForService(ctx, req.ServiceName)
	if err != nil {
		return errors.Wrap(err, "error getting service data dependencies")
	}

	oldDepsMap := make(map[string]domain.Resource)
	for _, d := range depsList {
		oldDepsMap[d.Name] = d
	}

	newDeps := make([]domain.Resource, 0, len(config.DataSources))

	for _, dep := range config.DataSources {
		_, ok := oldDepsMap[dep.GetName()]
		if ok {
			continue
		}

		depSrv := domain.Resource{
			Name:        req.ServiceName + "_" + dep.GetName(),
			ServiceName: req.ServiceName,
			Image:       patterns.GetImageNameByType(dep.GetType()),
			State:       domain.ServiceStateCreated,
		}

		if depSrv.Image == "" {
			continue
		}

		newDeps = append(newDeps, depSrv)
	}

	err = r.resourcesData.Upsert(ctx, newDeps...)
	if err != nil {
		return errors.Wrap(err, "error upserting dependencies")
	}

	return nil
}
