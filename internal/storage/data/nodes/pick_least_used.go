package nodes

import (
	"context"
	"database/sql"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) ListLeastUsedNodes(ctx context.Context, req domain.PickNodesReq) ([]domain.VelezConnection, error) {
	r, err := p.db.QueryContext(ctx, `
		SELECT 
		    node.name,
		    node.addr,
		    
		    node.velez_port,
		    node.custom_velez_key_path,
			node.is_insecure,
			
			node.ssh_key,
			node.ssh_port,
			node.ssh_user_name
		FROM nodes node
		LEFT JOIN instances inst ON inst.node_name = node.name
		
		GROUP BY node.name
		ORDER BY count(inst.service_name)
		LIMIT $1
`, req.NodesCount)
	if err != nil {
		return nil, errors.Wrap(err, "error getting nodes")
	}

	defer r.Close()

	out := make([]domain.VelezConnection, 0, req.NodesCount)
	for r.Next() {
		var v domain.VelezConnection
		v, err = scanVelezConnection(r)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning nodes")
		}

		out = append(out, v)
	}

	err = r.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error after scanning rows")
	}

	return out, nil
}

func scanVelezConnection(row *sql.Rows) (node domain.VelezConnection, err error) {
	var key *[]byte
	var username *string
	var port *uint64

	err = row.Scan(
		&node.Node.Name,
		&node.Node.Addr,
		&node.Node.Port,
		&node.Node.CustomVelezKeyPath,
		&node.Node.IsInsecure,

		&key,
		&port,
		&username,
	)
	if err != nil {
		return node, errors.Wrap(err, "error scanning")
	}

	if key != nil && port != nil && username != nil {
		node.Ssh = &domain.Ssh{
			Key:      *key,
			Port:     *port,
			Username: *username,
		}
	}

	return node, nil
}
