package nodes

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) SaveConnection(ctx context.Context, in domain.VelezConnection) error {
	tx, err := p.db.Begin()
	if err != nil {
		return errors.Wrap(err, "error starting transaction")
	}
	defer func() {
		// TODO move to pkg
		var txErr error
		if err != nil {
			txErr = tx.Rollback()
		} else {
			txErr = tx.Commit()
		}

		if txErr != nil {
			err = errors.Wrap(txErr, "error closing transaction")
		}
	}()

	_, err = p.db.ExecContext(ctx, `
		INSERT INTO nodes
				(name, addr, velez_port, custom_velez_key_path, is_insecure)
		VALUES  (  $1,   $2,         $3, 	                 $4,          $5)
		ON CONFLICT DO UPDATE SET
				 name 				    = excluded.name,   
				 velez_port 			= excluded.velez_port, 
				 addr 					= excluded.addr,
				 custom_velez_key_path 	= excluded.custom_velez_key_path, 
				 is_insecure 			= excluded.is_insecure
`,
		in.Node.Name, in.Node.Addr, in.Node.Port,
		in.Node.CustomVelezKeyPath, in.Node.IsInsecure,
	)
	if err != nil {
		return errors.Wrap(err, "error saving velez node")
	}
	if in.Ssh != nil {
		_, err = p.db.ExecContext(ctx, `
		UPDATE nodes
		SET ssh_key       = $1,
		    ssh_port      = $2,
		    ssh_user_name = $3 
     	WHERE name = $4
`, in.Ssh.Key, in.Ssh.Port, in.Ssh.Username, in.Node.Name)
		if err != nil {
			return errors.Wrap(err, "error saving ssh info")
		}
	}

	return nil
}
