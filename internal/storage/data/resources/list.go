package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
)

func (p *Provider) ListForService(ctx context.Context, name string) ([]domain.Resource, error) {
	row, err := p.conn.QueryContext(ctx, ``, name)
	if err != nil {
		return nil, errors.Wrap(err, "error querying database")
	}

	defer row.Close()

	out := make([]domain.Resource, 0)
	for row.Next() {
		var res domain.Resource
		res, err = toResource(row)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning db response")
		}

		out = append(out, res)
	}

	return out, nil
}

func toResource(row storage.Row) (r domain.Resource, err error) {
	return r, row.Scan(
		&r.Name,
		&r.ServiceName,
		&r.Image,
		&r.State,
	)
}
