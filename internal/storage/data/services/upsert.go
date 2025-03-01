package services

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (s *Provider) Upsert(ctx context.Context, services ...domain.Service) error {
	stmp, err := s.db.PrepareContext(ctx, `
		INSERT INTO services 
			   ( name, image, state, replicas)
		VALUES (   $1,    $2,    $3,       $4)
		ON CONFLICT (name) DO UPDATE SET 
			 name = excluded.name,
			 image = excluded.image, 
			 state = excluded.state,
			 replicas = excluded.replicas
		`)
	if err != nil {
		return errors.Wrap(err, "error creating prepare statement")
	}
	defer stmp.Close()

	for _, srv := range services {
		_, err = stmp.Exec(srv.Name, srv.Image, srv.State, srv.Replicas)
		if err != nil {
			return errors.Wrap(err, "error upserting service to storage")
		}
	}

	return nil
}
