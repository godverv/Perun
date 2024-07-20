package resource_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *ResourceService) GetDiffForService(ctx context.Context, serviceName string) (domain.ResourceDiff, error) {
	srv, err := s.servicesData.Get(ctx, serviceName)
	if err != nil {
		return domain.ResourceDiff{}, errors.Wrap(err, "error getting service by name")
	}

	registeredResources, err := s.resourcesData.ListForService(ctx, srv.Name)
	if err != nil {
		return domain.ResourceDiff{}, errors.Wrap(err, "error listing registered resources for service")
	}

	cfg, err := s.configService.FetchForService(ctx, srv)
	if err != nil {
		return domain.ResourceDiff{}, errors.Wrap(err, "error fetching registered resources for service")
	}

	diff := s.getDiff(registeredResources, fromConfig(serviceName, cfg))
	return diff, nil
}

func (s *ResourceService) getDiff(registered, fromConfig []domain.Resource) domain.ResourceDiff {
	registeredMap := make(map[string]domain.Resource)
	for _, r := range registered {
		registeredMap[r.Name] = r
	}

	var diff domain.ResourceDiff

	for _, cfgRes := range fromConfig {
		res, ok := registeredMap[cfgRes.Name]
		if !ok {
			diff.NewResources = append(diff.NewResources, cfgRes)
		} else {
			if res.State != domain.ServiceStateRunningOk {
				diff.StoppedResources = append(diff.StoppedResources, res)
			}
		}
	}

	return diff
}
