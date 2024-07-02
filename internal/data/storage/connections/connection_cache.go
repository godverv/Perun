package connections

import (
	"sync"

	"github.com/Red-Sock/Perun/internal/domain"
)

type ConnectionCache struct {
	mut     sync.RWMutex
	clients map[string]domain.Node
}

func NewConnectionCache() *ConnectionCache {
	return &ConnectionCache{
		clients: make(map[string]domain.Node),
	}
}

func (c *ConnectionCache) Add(clients ...domain.Node) {
	c.mut.Lock()
	for _, cl := range clients {
		c.clients[cl.Name] = cl
	}
	c.mut.Unlock()
}

func (c *ConnectionCache) Get(names ...string) ([]domain.Node, error) {
	c.mut.RLock()
	defer c.mut.RUnlock()

	out := make([]domain.Node, 0, len(c.clients))

	for _, name := range names {
		cl, ok := c.clients[name]
		_ = ok // todo
		out = append(out, cl)
	}

	return out, nil
}
