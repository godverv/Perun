package instances

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
)

func (p *Provider) List(ctx context.Context, req domain.ListInstancesReq) ([]domain.Instance, error) {
	selectQuery := sq.Select(
		"name",
		"node_name",
		"port",
		"state",
		"image_name").
		From("instances").
		PlaceholderFormat(sq.Dollar)

	if len(req.Names) != 0 {
		nameWhere := sq.Or{}
		for _, n := range req.Names {
			nameWhere = append(nameWhere, sq.Eq{"name": n})
		}
		selectQuery = selectQuery.Where(nameWhere)
	}

	sql, args, err := selectQuery.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "error building select")
	}

	row, err := p.conn.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error selecting instances")
	}
	defer row.Close()

	out := make([]domain.Instance, 0)
	for row.Next() {
		var ins domain.Instance
		ins, err = scanToInstance(row)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning ")
		}

		out = append(out, ins)
	}

	return out, nil
}

func scanToInstance(r storage.Row) (out domain.Instance, err error) {
	return out, r.Scan(
		&out.Name,
		&out.NodeName,
		&out.Port,
		&out.State,
		&out.Image,
	)
}
