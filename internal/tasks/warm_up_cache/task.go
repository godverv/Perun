package warm_up_cache

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/storage/connections_cache"
)

type WarmUpCacheTask struct {
	nodes storage.Nodes

	cache storage.ConnectionCache
}

func New(d storage.Data) WarmUpCacheTask {
	return WarmUpCacheTask{
		nodes: d.Nodes(),
		cache: d.Connections(),
	}
}

func (t WarmUpCacheTask) Do() {
	err := t.do()
	if err != nil {
		logrus.Errorf("error doing warm up cache task: %s", err.Error())
	}
}

func (t WarmUpCacheTask) do() error {
	ctx := context.Background()
	node, err := t.nodes.ListLeastUsedNodes(ctx, domain.PickNodeReq{ReplicationFactor: 10})
	if err != nil {
		return errors.Wrap(err, "error getting least used nodes")
	}

	for _, n := range node {
		conn, err := connections_cache.NewVelezService(n)
		if err != nil {
			return errors.Wrap(err, "error creating velez service")
		}
		t.cache.Add(domain.Node{
			Name: n.Node.Name,
			Conn: conn,
		})
	}

	return nil
}
