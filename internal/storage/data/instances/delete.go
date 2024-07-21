package instances

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
)

func (p *Provider) Delete(ctx context.Context, name string) error {
	_, err := p.conn.ExecContext(ctx, `
		DELETE FROM instances 
	   	WHERE name = $1
`, name)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
