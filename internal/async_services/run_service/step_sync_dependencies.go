package run_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

var ErrCreatedResourceHasNoPortsToAccess = errors.New("created resource has no ports to access")

type SyncDependenciesStep struct {
	resourceService service.ResourceService
	resourceData    storage.Resources

	configService matreshka_api.MatreshkaBeAPIClient
}

func NewSyncDependenciesStep(
	resourceService service.ResourceService,
	resourceData storage.Resources,
	configService matreshka_api.MatreshkaBeAPIClient,
) *SyncDependenciesStep {
	return &SyncDependenciesStep{
		resourceService: resourceService,
		resourceData:    resourceData,
		configService:   configService,
	}
}

func (s *SyncDependenciesStep) Do(ctx context.Context, r *RunServiceReq) error {
	if len(r.Nodes) == 0 {
		return errors.Wrap(ErrNoNodes, "no nodes provided to sync dependencies")
	}

	dependencies, err := s.resourceService.GetDependencies(ctx, r.ServiceName)
	if err != nil {
		return errors.Wrap(err, "error getting dependencies")
	}

	for _, dep := range dependencies.Resources {
		var res []domain.Resource
		res, err = s.resourceData.Get(ctx, dep.Name)
		if err != nil {
			return errors.Wrap(err, "error getting dependency from db")
		}

		if len(res) == 0 {
			r.Dependencies.Resources = append(r.Dependencies.Resources, dep)
		}
	}

	return nil
}
