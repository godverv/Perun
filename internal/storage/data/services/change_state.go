package services

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *Services) UpdateState(ctx context.Context, service domain.Service) error {
	_, err := s.db.ExecContext(ctx, `
		UPDATE services 
		SET 
			state = $1
		WHERE name = $2`,
		service.State,
		service.Name,
	)
	if err != nil {
		return errors.Wrap(err, "error performing sql update")
	}

	return nil
}
