package async_services

import (
	"context"

	"github.com/Red-Sock/Perun/internal/domain"
)

type EventHandler[E any] interface {
	Consume(ctx context.Context, event E) error
}

type ConsumerQueue[E any] interface {
	Dispatch(event E) error
}

type Queue[E any] interface {
	ConsumerQueue[E]

	Subscribe(handler EventHandler[E])
	Stop()
}

type AsyncService interface {
	InitServiceQueue() ConsumerQueue[domain.InitServiceReq]
	RefreshServiceQueue() ConsumerQueue[domain.RefreshService]
	DeployServiceQueue() ConsumerQueue[domain.DeployServiceReq]
}
