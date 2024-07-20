package init_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/async_services"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
)

type InitService struct {
	resourcesStorage storage.Services

	refreshServiceInfoQueue async_services.ConsumerQueue[domain.RefreshService]
}

func New(data storage.Data, asyncSrv async_services.AsyncService) *InitService {
	return &InitService{
		resourcesStorage:        data.Services(),
		refreshServiceInfoQueue: asyncSrv.RefreshServiceQueue(),
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

	refreshServiceEvent := domain.RefreshService{
		ServiceName: newService.Name,
	}
	err = s.refreshServiceInfoQueue.Dispatch(refreshServiceEvent)
	if err != nil {
		return errors.Wrap(err, "error dispatching refresh service info event")
	}

	return nil
}
