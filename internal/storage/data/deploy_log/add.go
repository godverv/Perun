package deploy_log

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) Add(ctx context.Context, in domain.DeployLog) error {
	_, err := p.conn.ExecContext(ctx, `
			INSERT INTO deploy_log
					(name, state, reason) 
			VALUES  (  $1,    $2,     $3) 
`,
		in.Name, in.State, in.Reason)
	if err != nil {
		return errors.Wrap(err, "error saving deploy log")
	}

	return nil
}
