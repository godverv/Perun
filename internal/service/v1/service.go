package v1

import (
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/service/v1/config_service"
	"github.com/Red-Sock/Perun/internal/service/v1/deploy_log"
	"github.com/Red-Sock/Perun/internal/service/v1/nodes_service"
	"github.com/Red-Sock/Perun/internal/service/v1/resource_service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type services struct {
	nodes           service.NodesService
	configService   service.ConfigService
	resourceService service.ResourceService
	deployLog       service.DeployLog
}

func NewService(storage storage.Data) service.Services {
	nodeSrv := nodes_service.New(storage)
	cfgSrv := config_service.New(nodeSrv)
	return &services{
		nodes:           nodeSrv,
		configService:   cfgSrv,
		resourceService: resource_service.New(storage, nodeSrv, cfgSrv),
		deployLog:       deploy_log.New(storage),
	}
}

func (s *services) Nodes() service.NodesService {
	return s.nodes
}

func (s *services) Config() service.ConfigService {
	return s.configService
}

func (s *services) Resources() service.ResourceService {
	return s.resourceService
}

func (s *services) DeployLog() service.DeployLog {
	return s.deployLog
}
