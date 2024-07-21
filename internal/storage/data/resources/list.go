package resources

import (
	"context"
	"strconv"
	"strings"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) List(ctx context.Context, req domain.DeployResourcesReq) ([]domain.Resource, error) {
	sb := strings.Builder{}
	sb.WriteString(`
			SELECT 
			    name,
				service_name,
				image,
				state
			FROM resources
			WHERE name IN (		
`)

	var args []any
	for idx := range req.ResourcesNames {
		sb.WriteByte('$')
		sb.WriteString(strconv.Itoa(idx + 1))
		if idx < len(req.ResourcesNames)-1 {
			sb.WriteByte(',')
		}
		args = append(args, req.ResourcesNames[idx])
	}
	sb.WriteByte(')')
	q := sb.String()
	rows, err := p.conn.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error listing resources")
	}
	defer rows.Close()

	out := make([]domain.Resource, 0)
	for rows.Next() {
		var r domain.Resource
		r, err = toResource(rows)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning")
		}

		out = append(out, r)
	}

	return out, nil
}
