package instances

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) Update(ctx context.Context, instances domain.Instance) error {
	_, err := p.conn.ExecContext(ctx, `
			UPDATE instances
			SET 
			    node_name  = $1,
				port       = $2,
				state      = $3,
				image_name = $4
			WHERE name = $5 
`,
		instances.NodeName,
		instances.Port,
		instances.State,
		instances.Image,
		instances.Name)
	if err != nil {
		return errors.Wrap(err, "error performing update")
	}

	return nil
}
