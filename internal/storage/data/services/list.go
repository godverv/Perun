package services

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
)

func (s *Provider) List(ctx context.Context, serviceNamePattern string) ([]domain.Service, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT 
			name,
			image,
			state,
			replicas
		FROM services
		WHERE name like '%'||$1||'%'
`, serviceNamePattern)
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

func toService(row storage.Row) (res domain.Service, err error) {
	return res, row.Scan(
		&res.Name,
		&res.Image,
		&res.State,
		&res.Replicas,
	)
}
