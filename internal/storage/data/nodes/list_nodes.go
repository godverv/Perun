package nodes

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

const (
	nodesLimitMax = 100
	nodesLimitMin = 10
)

func (p *Provider) ListNodes(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConnection, error) {
	if req.Limit > nodesLimitMax {
		req.Limit = nodesLimitMax
	}
	if req.Limit == 0 {
		req.Limit = nodesLimitMin
	}

	row, err := p.db.QueryContext(ctx, `
		SELECT 
		    n.node_name,
		    n.addr,
		    n.velez_port,
		    n.custom_velez_key_path,
			n.is_insecure,
			
			n.ssh_key,
			n.ssh_port,
			n.ssh_user_name
		FROM nodes n
		WHERE node_name LIKE '%'||$1||'%'
		LIMIT $2 OFFSET $3
`, req.SearchPattern, req.Limit, req.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "error listing velez nodes")
	}
	defer row.Close()

	nodes := make([]domain.VelezConnection, 0, req.Limit)

	for row.Next() {
		var v domain.VelezConnection
		v, err = scanVelezConnection(row)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning velez connection")
		}
		nodes = append(nodes, v)
	}

	err = row.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error scanning rows")
	}

	return nodes, nil
}
