package v1

import (
	"github.com/godverv/matreshka-be/pkg/matreshka_api"

	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/service/v1/nodes_service"
	"github.com/Red-Sock/Perun/internal/service/v1/resource_service"
)

type services struct {
	nodes     service.NodesService
	resources service.ResourceService
}

func NewService(storage data.Data, client matreshka_api.MatreshkaBeAPIClient) service.Services {
	nodes := nodes_service.New(storage)

	resources := resource_service.New(client)

	return &services{
		nodes:     nodes,
		resources: resources,
	}
}

func (s *services) Resources() service.ResourceService {
	return s.resources
}

func (s *services) Nodes() service.NodesService {
	return s.nodes
}
