package sync_service_info

import (
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service/v1/constructors/patterns"
)

func (r *RefreshServiceConfig) collectDependencies(srvName string, cfg matreshka.AppConfig) (domain.Dependencies, error) {
	var deps domain.Dependencies

	for _, cfgResource := range cfg.DataSources {
		constructor := patterns.GetConstructor(cfgResource.GetType())
		if constructor == nil {
			continue
		}

		resourceDeps, err := constructor(cfg.DataSources, cfgResource.GetName(), srvName)
		if err != nil {
			return deps, errors.Wrap(err, "error getting resource-smerd config ")
		}

		deps.Resources = append(deps.Resources, resourceDeps.Resources...)
	}

	return deps, nil
}
