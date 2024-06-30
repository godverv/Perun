package v1

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"

	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/service/v1/nodes_service"
	"github.com/Red-Sock/Perun/internal/service/v1/resource_service"
	"github.com/Red-Sock/Perun/internal/service/v1/service_runner"
)

type services struct {
	nodes     service.NodesService
	resources service.ResourceService

	runner service.RunnerService
}

func NewService(ctx context.Context, storage data.Storage, client matreshka_api.MatreshkaBeAPIClient) (service.Services, error) {
	nodes, err := nodes_service.New(ctx, storage)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing nodes_service")
	}
	resources := resource_service.New(client)

	return &services{
		nodes:     nodes,
		resources: resources,
		runner:    service_runner.New(resources, storage),
	}, nil
}

func (s *services) Resources() service.ResourceService {
	return s.resources
}

func (s *services) Runner() service.RunnerService {
	return s.runner
}

func (s *services) Nodes() service.NodesService {
	return s.nodes
}
