package init_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
)

type InitService struct {
	resourcesStorage storage.Services
}

func New(data storage.Data) *InitService {
	return &InitService{
		resourcesStorage: data.Services(),
	}
}

func (s *InitService) Consume(ctx context.Context, req domain.InitServiceReq) error {
	newService := domain.Service{
		Name:     req.ServiceName,
		Image:    req.ImageName,
		State:    domain.ServiceStateCreated,
		Replicas: req.ReplicationFactor,
	}

	err := s.resourcesStorage.Upsert(ctx, newService)
	if err != nil {
		return errors.Wrap(err, "error upserting service")
	}

	return nil
}
