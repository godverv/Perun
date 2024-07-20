package resource_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service/v1/constructors/patterns"
)

func (s *ResourceService) GetFromConfig(ctx context.Context, serviceName string) ([]domain.Resource, error) {
	srv, err := s.servicesData.Get(ctx, serviceName)
	if err != nil {
		return nil, errors.Wrap(err, "error getting service")
	}

	cfg, err := s.configService.FetchForService(ctx, srv)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching config")
	}

	return fromConfig(serviceName, cfg), nil
}

func fromConfig(serviceName string, cfg matreshka.AppConfig) []domain.Resource {
	resources := make([]domain.Resource, 0, len(cfg.DataSources))

	for _, dep := range cfg.DataSources {
		image := patterns.GetImageNameByType(dep.GetType())
		if image == "" {
			continue
		}

		depSrv := domain.Resource{
			Name:        serviceName + "_" + dep.GetName(),
			ServiceName: serviceName,
			Image:       image,
			State:       domain.ServiceStateCreated,
		}

		resources = append(resources, depSrv)
	}

	return resources
}
