package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Storage) UpdateState(ctx context.Context, changeState domain.UpdateState) error {
	_, err := p.db.ExecContext(ctx, `
		UPDATE resources 
		SET state = $1
		WHERE resource_full_name = $2`, changeState.State, changeState.ResourceName)
	if err != nil {
		return errors.Wrap(err, "error performing sql update")
	}

	return nil
}
