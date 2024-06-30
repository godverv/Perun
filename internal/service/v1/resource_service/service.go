package resource_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service/v1/resource_service/resource_patterns"
)

type ResourceManager struct {
	matreshkaApi matreshka_api.MatreshkaBeAPIClient
}

func New(matreshkaApi matreshka_api.MatreshkaBeAPIClient) *ResourceManager {
	return &ResourceManager{
		matreshkaApi: matreshkaApi,
	}
}

func (m *ResourceManager) GetDependencies(ctx context.Context, serviceName string) (domain.Dependencies, error) {
	cfg, err := m.getConfig(ctx, serviceName)
	if err != nil {
		return domain.Dependencies{}, errors.Wrap(err, "error getting config")
	}

	var deps domain.Dependencies
	for _, cfgResource := range cfg.DataSources {
		constructor := resource_patterns.GetConstructor(cfgResource.GetType())
		if constructor == nil {
			continue
		}

		resourceDeps, err := constructor(cfg.DataSources, cfgResource.GetName(), serviceName)
		if err != nil {
			return deps, errors.Wrap(err, "error getting resource-smerd config ")
		}

		deps.Volumes = append(deps.Volumes, resourceDeps.Volumes...)
		deps.Smerds = append(deps.Smerds, resourceDeps.Smerds...)
	}

	return deps, nil
}

func (m *ResourceManager) GetConfig(ctx context.Context, serviceName string) (matreshka.AppConfig, error) {
	return m.getConfig(ctx, serviceName)
}

func (m *ResourceManager) getConfig(ctx context.Context, serviceName string) (matreshka.AppConfig, error) {
	getConfigReq := &matreshka_api.GetConfig_Request{ServiceName: serviceName}
	cfgResp, err := m.matreshkaApi.GetConfig(ctx, getConfigReq)
	if err != nil {
		return matreshka.AppConfig{}, errors.Wrap(err, "error getting config")
	}

	cfg := matreshka.NewEmptyConfig()
	err = cfg.Unmarshal(cfgResp.Config)
	if err != nil {
		return matreshka.AppConfig{}, errors.Wrap(err, "error unmarshalling config to struct")
	}

	return cfg, nil
}
