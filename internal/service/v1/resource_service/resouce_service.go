package resource_service

import (
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type ResourceService struct {
	servicesData  storage.Services
	resourcesData storage.Resources

	nodeService   service.NodesService
	configService service.ConfigService
}

func New(data storage.Data, nodeService service.NodesService, configService service.ConfigService) *ResourceService {
	return &ResourceService{
		servicesData:  data.Services(),
		resourcesData: data.Resources(),

		nodeService:   nodeService,
		configService: configService,
	}
}
