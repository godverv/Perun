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
			ssh_key,
			ssh_addr,
			ssh_user_name,
			velez_addr,
			custom_velez_key_path,
			insecure
		FROM nodes
		WHERE node_name = $1
    `, nodeName).
		Scan(
			&v.Node.Name,
			&v.Ssh.Key,
			&v.Ssh.Addr,
			&v.Ssh.Username,
			&v.Node.Addr,
			&v.Node.CustomVelezKeyPath,
			&v.Node.IsInsecure,
		)
	if err != nil {
		return nil, errors.Wrap(err, "error getting velez node info")
	}

	return nil, nil
}
