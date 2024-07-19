package in_async_service

import (
	"github.com/Red-Sock/Perun/internal/async_services"
	"github.com/Red-Sock/Perun/internal/async_services/in_async_service/publishers"
	"github.com/Red-Sock/Perun/internal/async_services/in_async_service/subscribers/init_service"
	"github.com/Red-Sock/Perun/internal/async_services/in_async_service/subscribers/sync_service_info"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type Queue struct {
	initServiceQueue     async_services.Queue[domain.InitServiceReq]
	syncServiceInfoQueue async_services.Queue[domain.SyncServiceInfo]
}

func New(data storage.Data, srv service.Services) *Queue {
	q := &Queue{
		initServiceQueue:     publishers.New[domain.InitServiceReq](),
		syncServiceInfoQueue: publishers.New[domain.SyncServiceInfo](),
	}

	q.initServiceQueue.Subscribe(init_service.New(data))

	q.syncServiceInfoQueue.Subscribe(sync_service_info.New(data, srv))

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

func (q *Queue) RefreshServiceQueue() async_services.ConsumerQueue[domain.SyncServiceInfo] {
	return q.syncServiceInfoQueue
}
