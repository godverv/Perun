package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *Storage) List(ctx context.Context, req domain.ListResources) ([]domain.Resource, error) {
	out := make([]domain.Resource, 0)

	rows, err := s.db.QueryContext(ctx, `
		SELECT 
			resource_full_name,
			node_name, 
			state,
			port
		FROM resources
		WHERE resource_full_name like '%'||$1||'%'
`, req.Name)
	if err != nil {
		return nil, errors.Wrap(err, "error listing resources")
	}
	defer rows.Close()

	for rows.Next() {
		var res domain.Resource
		err = rows.Scan(
			&res.ResourceName,
			&res.NodeName,
			&res.State,
			&res.Port,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning row")
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error scanning resources")
	}
	return out, nil
}
