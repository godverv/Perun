package nodes_service

import (
	"github.com/Red-Sock/Perun/internal/storage"
)

type NodesService struct {
	nodesStore storage.Nodes

	connectionsCache storage.ConnectionCache
}

func New(store storage.Data) *NodesService {
	return &NodesService{
		nodesStore:       store.Nodes(),
		connectionsCache: store.Connections(),
	}
}
