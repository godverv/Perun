package service

import (
	"context"

	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
)

type Services interface {
	Nodes() NodesService
	Config() ConfigService
	Resources() ResourceService
}

type NodesService interface {
	NewVelezConnection(ctx context.Context, node domain.VelezConnection) error

	ListNodes(ctx context.Context, nodes domain.ListVelezNodes) ([]domain.VelezConnection, error)
	PickNodes(ctx context.Context, req domain.PickNodesReq) ([]domain.Node, error)
}

type ConfigService interface {
	FetchForService(ctx context.Context, srv domain.Service) (matreshka.AppConfig, error)
}

type ResourceService interface {
	GetDiffForService(ctx context.Context, serviceName string) (domain.ResourceDiff, error)
}
