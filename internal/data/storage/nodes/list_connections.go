package nodes

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) ListConnections(ctx context.Context, req domain.ListVelezNodes) ([]domain.VelezConnection, error) {
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
		    insecure,
		    
		    ssh_key,
		    ssh_addr,
		    ssh_user_name
		FROM nodes
		WHERE node_name LIKE '%'||$1||'%'
		LIMIT $2 OFFSET $3
`, req.SearchPattern, req.Limit, req.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "error listing velez nodes")
	}
	defer row.Close()

	nodes := make([]domain.VelezConnection, 0, req.Limit)

	for row.Next() {
		var conn domain.VelezConnection
		err = row.Scan(
			&conn.Node.Name,
			&conn.Node.Addr,
			&conn.Node.CustomVelezKeyPath,
			&conn.Node.IsInsecure,

			&conn.Ssh.Key,
			&conn.Ssh.Addr,
			&conn.Ssh.Username,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning row")
		}

		nodes = append(nodes, conn)
	}

	err = row.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error scanning rows")
	}

	return nodes, nil
}
