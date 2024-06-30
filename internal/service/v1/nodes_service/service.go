package nodes_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/data/velez"
	"github.com/Red-Sock/Perun/internal/domain"
)

type nodesService struct {
	nodes      map[string]domain.Node
	filedNodes map[string]domain.Node
	nodesStore data.Nodes
}

func New(ctx context.Context, store data.Storage) (*nodesService, error) {
	ns := &nodesService{
		nodesStore: store.Nodes(),
		nodes:      make(map[string]domain.Node),
		filedNodes: make(map[string]domain.Node),
	}

	connections, err := ns.nodesStore.ListConnections(ctx, domain.ListVelezNodes{})
	if err != nil {
		return nil, errors.Wrap(err, "error listing connections")
	}

	for _, connSetup := range connections {
		node := domain.Node{
			Name: connSetup.Node.Name,
		}

		node.Conn, err = velez.NewVelezService(connSetup)
		if err != nil {
			ns.filedNodes[node.Name] = node
			logrus.Warning("error creating velez service connection: " + err.Error())
		} else {
			ns.nodes[node.Name] = node
		}
	}

	return ns, nil
}
