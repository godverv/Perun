package sync_service_info

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service/v1/constructors/patterns"
)

func (r *RefreshServiceConfig) syncDependencies(ctx context.Context,
	req domain.SyncServiceInfo,
	config matreshka.AppConfig,
) error {
	depsList, err := r.servicesData.ListChildren(ctx, req.Service.Name)
	if err != nil {
		return errors.Wrap(err, "error getting service data dependencies")
	}
	oldDepsMap := make(map[string]domain.Service)
	for _, d := range depsList {
		oldDepsMap[d.Name] = d
	}

	newDeps := make([]domain.Service, 0, len(config.DataSources))

	for _, dep := range config.DataSources {
		_, ok := oldDepsMap[dep.GetName()]
		if ok {
			continue
		}

		depSrv := domain.Service{
			Name:  req.Service.Name + "_" + dep.GetName(),
			Image: patterns.GetImageNameByType(dep.GetType()),
			State: domain.ServiceStateCreated,
		}

		newDeps = append(newDeps, depSrv)
	}

	err = r.servicesData.Upsert(ctx, newDeps...)
	if err != nil {
		return errors.Wrap(err, "error upserting dependencies")
	}

	return nil
}
