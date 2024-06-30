package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Storage) Create(ctx context.Context, resource domain.Resource) error {
	_, err := p.db.ExecContext(ctx, `
		INSERT INTO resources 
		 	   (resource_full_name, node_name, state)
		VALUES (                $1,        $2,    $3)`,
		resource.ResourceName, resource.NodeName, resource.State)
	if err != nil {
		return errors.Wrap(err, "error creating resource")
	}
	return nil
}
