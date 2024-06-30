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

func (p *Provider) ListNodes(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConn, error) {
	if req.Limit > nodesLimitMax {
		req.Limit = nodesLimitMax
	}
	if req.Limit == 0 {
		req.Limit = nodesLimitMin
	}

	row, err := p.db.QueryContext(ctx, `
		SELECT 
		    node_name,
		    velez_addr,
		    custom_velez_key_path,
		    insecure
		FROM nodes
		WHERE node_name LIKE '%'||$1||'%'
		LIMIT $2 OFFSET $3
`, req.SearchPattern, req.Limit, req.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "error listing velez nodes")
	}
	defer row.Close()

	nodes := make([]domain.VelezConn, 0, req.Limit)

	for row.Next() {
		var node domain.VelezConn
		err = row.Scan(
			&node.Name,
			&node.Addr,
			&node.CustomVelezKeyPath,
			&node.IsInsecure,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning row")
		}

		nodes = append(nodes, node)
	}

	err = row.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error scanning rows")
	}

	return nodes, nil
}
