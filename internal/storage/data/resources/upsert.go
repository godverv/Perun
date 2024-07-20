package resources

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (p *Provider) Upsert(ctx context.Context, deps ...domain.Resource) error {
	stmp, err := p.conn.PrepareContext(ctx, `
		INSERT INTO resources
		    	(name, service_name, image, state) 
		VALUES  (  $1,           $2,    $3,    $4)
		ON CONFLICT (name)
		DO UPDATE SET
			service_name = excluded.service_name,
			image 		 = excluded.image,
			state 		 = excluded.state	 
`)
	if err != nil {
		return errors.Wrap(err, "error creating prepare statement")
	}
	defer stmp.Close()

	for _, d := range deps {
		_, err = stmp.Exec(d.Name, d.ServiceName, d.Image, d.State)
		if err != nil {
			return errors.Wrap(err, "error performing upsert statement")
		}
	}

	return nil
}
