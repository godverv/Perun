package services

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *Provider) Get(ctx context.Context, name string) (domain.Service, error) {
	scanner := s.db.QueryRowContext(ctx, `
		SELECT 
			name,
			image,
			state,
			replicas
		FROM services
		WHERE name = $1
`, name)

	service, err := toService(scanner)
	if err != nil {
		return domain.Service{}, errors.Wrap(err, "error listing resources")
	}

	return service, nil
}
