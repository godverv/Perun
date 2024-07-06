package nodes

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) GetConnection(ctx context.Context, nodeName string) (*domain.VelezConnection, error) {
	var v domain.VelezConnection

	err := p.db.QueryRowContext(ctx, `
		SELECT
		   	node_name,
			addr,
			velez_port,
			custom_velez_key_path,
			is_insecure,
			
			ssh_key,
			ssh_port,
			ssh_user_name
			
		FROM nodes
		WHERE node_name = $1
    `, nodeName).
		Scan(
			&v.Node.Name,
			&v.Node.Addr,
			&v.Node.Port,
			&v.Node.CustomVelezKeyPath,
			&v.Node.IsInsecure,

			&v.Ssh.Key,
			&v.Ssh.Port,
			&v.Ssh.Username,
		)
	if err != nil {
		return nil, errors.Wrap(err, "error getting velez node info")
	}

	return nil, nil
}
