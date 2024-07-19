package services

import (
	"context"
	"database/sql"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *Services) Get(ctx context.Context, name string) (*domain.Service, error) {
	scanner := s.db.QueryRowContext(ctx, `
		SELECT 
			name,
			image,
			state
		FROM services
		WHERE name = $1
`, name)

	service, err := toService(scanner)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "error listing resources")
	}

	return &service, nil
}
