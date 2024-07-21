package instances

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) Add(ctx context.Context, instance domain.Instance) error {
	_, err := p.conn.ExecContext(ctx, `
			INSERT INTO instances 
			  	   (name, node_name, port, state, image_name)
			VALUES (  $1,        $2,   $3,    $4,         $5)
`,
		instance.Name,
		instance.NodeName,
		instance.Port,
		instance.State,
		instance.Image)
	if err != nil {
		return errors.Wrap(err, "error inserting instance")
	}

	return nil
}
