package in_memory

import (
	"github.com/Red-Sock/Perun/internal/async_services"
	"github.com/Red-Sock/Perun/internal/async_services/in_memory/publishers"
	"github.com/Red-Sock/Perun/internal/async_services/in_memory/subscribers/deploy_service"
	"github.com/Red-Sock/Perun/internal/async_services/in_memory/subscribers/init_service"
	"github.com/Red-Sock/Perun/internal/async_services/in_memory/subscribers/refresh_service_config"
	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type Queue struct {
	initServiceQueue     async_services.Queue[domain.InitServiceReq]
	syncServiceInfoQueue async_services.Queue[domain.RefreshService]

	deployServiceQueue async_services.Queue[domain.DeployServiceReq]
}

func New(data storage.Data, srv service.Services) *Queue {
	q := &Queue{
		initServiceQueue:     publishers.New[domain.InitServiceReq](),
		syncServiceInfoQueue: publishers.New[domain.RefreshService](),
		deployServiceQueue:   publishers.New[domain.DeployServiceReq](),
	}

	q.initServiceQueue.Subscribe(init_service.New(data, q))

	q.syncServiceInfoQueue.Subscribe(refresh_service_config.New(data, srv))

	q.deployServiceQueue.Subscribe(deploy_service.New(data, srv))
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

func (q *Queue) DeployServiceQueue() async_services.ConsumerQueue[domain.DeployServiceReq] {
	return q.deployServiceQueue
}
