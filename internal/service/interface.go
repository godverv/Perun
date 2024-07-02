package service

import (
	"context"

	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
)

type Services interface {
	Nodes() NodesService
	Resources() ResourceService

	Runner() RunnerService
}

type NodesService interface {
	NewVelezConnection(ctx context.Context, node domain.VelezConnection) error
	ListNodes(ctx context.Context, nodes domain.ListVelezNodes) ([]domain.VelezConnection, error)
	PickNodes(ctx context.Context, req domain.PickNodeReq) ([]domain.Node, error)
}

type RunnerService interface {
	Run(ctx context.Context, node domain.RunServiceRequest) error
}

type ResourceService interface {
	GetDependencies(ctx context.Context, serviceName string) (domain.Dependencies, error)
	GetConfig(ctx context.Context, serviceName string) (matreshka.AppConfig, error)
}
