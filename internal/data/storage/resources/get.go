package resources

import (
	"context"
	"database/sql"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Storage) Get(ctx context.Context, name string) (res *domain.Resource, err error) {
	res = &domain.Resource{}

	res.ResourceName = name

	err = p.db.QueryRowContext(ctx,
		`
			SELECT
			    node_name
			FROM resources
			WHERE resource_full_name = $1
`, name).
		Scan(&res.NodeName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return res, errors.Wrap(err, "error reading resource from db")
	}

	return res, nil
}
