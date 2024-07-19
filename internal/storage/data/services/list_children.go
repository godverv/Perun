package services

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *Services) ListChildren(ctx context.Context, parentName string) ([]domain.Service, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT 
			name,
			image,
			state
		FROM services
		WHERE parent = $1
`, parentName)
	if err != nil {
		return nil, errors.Wrap(err, "error listing resources")
	}

	defer rows.Close()

	out := make([]domain.Service, 0)

	for rows.Next() {
		var row domain.Service
		row, err = toService(rows)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning row")
		}

		out = append(out, row)
	}

	return out, nil
}
