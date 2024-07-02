package nodes

import (
	"context"
	"database/sql"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) ListLeastUsedNodes(ctx context.Context, req domain.PickNodeReq) ([]domain.VelezConnection, error) {
	r, err := p.db.QueryContext(ctx, `
		SELECT 
		    n.node_name,
		    n.velez_addr,
		    n.custom_velez_key_path,
			n.insecure,
			
			n.ssh_addr,
			n.ssh_key,
			n.ssh_user_name
		FROM nodes n
		LEFT JOIN resources r ON r.node_name = n.node_name
		GROUP BY n.node_name
		ORDER BY count(r.resource_full_name)
		LIMIT $1
`, req.ReplicationFactor)
	if err != nil {
		return nil, errors.Wrap(err, "error getting nodes")
	}

	defer r.Close()

	out := make([]domain.VelezConnection, 0, req.ReplicationFactor)
	for r.Next() {
		out = append(out, scanVelezConnection(r))
	}

	err = r.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error after scanning rows")
	}

	return out, nil
}

func scanVelezConnection(row *sql.Rows) (node domain.VelezConnection) {
	_ = row.Scan(
		&node.Node.Name,
		&node.Node.Addr,
		&node.Node.CustomVelezKeyPath,
		&node.Node.IsInsecure,

		&node.Ssh.Addr,
		&node.Ssh.Key,
		&node.Ssh.Username,
	)

	return node
}
