package nodes

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) SaveConnection(ctx context.Context, in domain.VelezConnection) error {
	_, err := p.db.ExecContext(ctx, `
INSERT INTO nodes
		(node_name, addr, velez_port, custom_velez_key_path, is_insecure, ssh_key, ssh_port, ssh_user_name)
VALUES  (       $1,   $2,         $3, 	                 $4,          $5,      $6,       $7,            $8)
ON CONFLICT DO UPDATE SET
		 node_name 				= excluded.node_name,   
		 velez_port 			= excluded.velez_port, 
		 addr 					= excluded.addr,
		 custom_velez_key_path 	= excluded.custom_velez_key_path, 
		 is_insecure 			= excluded.is_insecure,
		 ssh_key 				= excluded.ssh_key,
		 ssh_port 				= excluded.ssh_port,
		 ssh_user_name 			= excluded.ssh_user_name
`,
		in.Node.Name, in.Node.Addr, in.Node.Port, in.Node.CustomVelezKeyPath, in.Node.IsInsecure,
		in.Ssh.Key, in.Ssh.Port, in.Ssh.Username)
	if err != nil {
		return errors.Wrap(err, "error saving velez node")
	}

	return nil
}
