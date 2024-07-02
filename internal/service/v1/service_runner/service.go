package service_runner

import (
	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/service"
)

type serviceRunner struct {
	resourceService service.ResourceService
	resourceData    data.Resources
}

func New(resourceService service.ResourceService, resourceData data.Data) *serviceRunner {
	p := &serviceRunner{
		resourceService: resourceService,
		resourceData:    resourceData.Resources(),
	}

	return p
}
