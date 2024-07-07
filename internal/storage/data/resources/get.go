package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Storage) Get(ctx context.Context, name string) ([]domain.Resource, error) {
	var out []domain.Resource

	rows, err := p.db.QueryContext(ctx,
		`
			SELECT
			    node_name, 
			    state,
			    port
			FROM resources
			WHERE resource_full_name = $1
`, name)
	if err != nil {
		return nil, errors.Wrap(err, "error reading resource from db")
	}
	defer rows.Close()

	for rows.Next() {
		var res domain.Resource
		res.ResourceName = name

		err = rows.Scan(
			&res.NodeName,
			&res.State,
			&res.Port,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error reading resource from rows")
		}

		out = append(out, res)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error reading rows from db")
	}

	return out, nil
}
