package v1

import (
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/service/v1/nodes_service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type services struct {
	nodes service.NodesService
}

func NewService(storage storage.Data) service.Services {
	return &services{
		nodes: nodes_service.New(storage),
	}
}

func (s *services) Nodes() service.NodesService {
	return s.nodes
}
