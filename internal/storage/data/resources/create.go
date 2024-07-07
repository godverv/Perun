package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *Storage) Create(ctx context.Context, resource domain.Resource) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO resources 
		 	   (resource_full_name, node_name, state, port)
		VALUES (                $1,        $2,    $3,   $4)`,
		resource.ResourceName, resource.NodeName, resource.State, resource.Port)
	if err != nil {
		return errors.Wrap(err, "error creating resource")
	}
	return nil
}
