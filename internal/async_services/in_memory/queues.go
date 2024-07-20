package in_memory

import (
	"github.com/Red-Sock/Perun/internal/async_services"
	"github.com/Red-Sock/Perun/internal/async_services/in_memory/publishers"
	"github.com/Red-Sock/Perun/internal/async_services/in_memory/subscribers/init_service"
	"github.com/Red-Sock/Perun/internal/async_services/in_memory/subscribers/refresh_service_config"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type Queue struct {
	initServiceQueue     async_services.Queue[domain.InitServiceReq]
	syncServiceInfoQueue async_services.Queue[domain.RefreshService]
}

func New(data storage.Data, srv service.Services) *Queue {
	q := &Queue{
		initServiceQueue:     publishers.New[domain.InitServiceReq](),
		syncServiceInfoQueue: publishers.New[domain.RefreshService](),
	}

	q.initServiceQueue.Subscribe(init_service.New(data, q))

	q.syncServiceInfoQueue.Subscribe(refresh_service_config.New(data, srv))

	return q
}

func (q *Queue) Stop() error {
	q.initServiceQueue.Stop()
	q.syncServiceInfoQueue.Stop()

	return nil
}

func (q *Queue) InitServiceQueue() async_services.ConsumerQueue[domain.InitServiceReq] {
	return q.initServiceQueue
}

func (q *Queue) RefreshServiceQueue() async_services.ConsumerQueue[domain.RefreshService] {
	return q.syncServiceInfoQueue
}
