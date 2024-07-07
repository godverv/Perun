package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Storage) Update(ctx context.Context, req domain.Resource) error {
	_, err := p.db.ExecContext(ctx, `
		UPDATE resources 
		SET 
			node_name = $1,
			state = $2,
			port = $3
		WHERE resource_full_name = $4`,
		req.NodeName,
		req.State,
		req.Port,
		req.ResourceName,
	)
	if err != nil {
		return errors.Wrap(err, "error performing sql update")
	}

	return nil
}
