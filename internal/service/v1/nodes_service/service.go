package nodes_service

import (
	"github.com/Red-Sock/Perun/internal/data"
)

type NodesService struct {
	nodesStore data.Nodes

	connectionsCache data.ConnectionCache
}

func New(store data.Data) *NodesService {
	return &NodesService{
		nodesStore:       store.Nodes(),
		connectionsCache: store.Connections(),
	}
}
