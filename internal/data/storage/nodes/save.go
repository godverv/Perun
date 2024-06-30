package nodes

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) SaveConnection(ctx context.Context, in domain.VelezConnection) error {
	_, err := p.db.ExecContext(ctx, `
INSERT INTO nodes
		(    node_name,    ssh_key,              ssh_addr,
		 ssh_user_name, velez_addr, custom_velez_key_path, insecure)
VALUES  (           $1,         $2,                    $3,
					$4,         $5,                    $6,       $7)
ON CONFLICT DO UPDATE SET
		 node_name 				= excluded.node_name,   
		 ssh_key 				= excluded.ssh_key,
		 ssh_addr 				= excluded.ssh_addr,
		 ssh_user_name 			= excluded.ssh_user_name, 
		 velez_addr 			= excluded.velez_addr,
		 custom_velez_key_path 	= excluded.custom_velez_key_path, 
		 insecure 				= excluded.insecure
`,
		in.Node.Name, in.Ssh.Key, in.Ssh.Addr, in.Ssh.Username,
		in.Node.Addr, in.Node.CustomVelezKeyPath, in.Node.IsInsecure)
	if err != nil {
		return errors.Wrap(err, "error saving velez node")
	}

	return nil
}
