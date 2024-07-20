package run_service

import (
	"context"
	"time"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
)

var defaultServiceUpWaitTime = time.Second * 10

type PostRunStep struct {
	resourcesStorage storage.Services
}

func NewPostRunStep(resourcesStorage storage.Services) *PostRunStep {
	return &PostRunStep{resourcesStorage: resourcesStorage}
}

func (p *PostRunStep) Do(ctx context.Context, r *RunServiceReq) error {
	time.Sleep(max(r.Config.StartupDuration, defaultServiceUpWaitTime))

	listReq := &velez_api.ListSmerds_Request{}
	listReq.Name = &r.ServiceName

	services := make([]*domain.Service, 0, len(r.Nodes))
	var err error
	defer func() {
		for _, service := range services {
			err = p.resourcesStorage.UpdateState(ctx, *service)
			if err != nil {
				err = errors.Wrap(err, "error updating resource state in storage")
				return
			}
		}
	}()

	for _, node := range r.Nodes {
		service := &domain.Service{
			Name: r.ServiceName,
			//NodeName: node.Name, TODO
		}

		services = append(services, service)
		var list *velez_api.ListSmerds_Response
		list, err = node.Conn.ListSmerds(ctx, listReq)
		if err != nil {
			return errors.Wrap(err, "error getting information from node")
		}

		var serviceContainer *velez_api.Smerd
		for _, smerd := range list.GetSmerds() {
			if smerd.GetName() == r.ServiceName {
				serviceContainer = smerd
				break
			}
		}
		if serviceContainer == nil {
			service.State = domain.ServiceStateError
			continue
		}

		if serviceContainer.Status != velez_api.Smerd_running {
			service.State = domain.ServiceStateError
			continue
		}

		if len(serviceContainer.Ports) != 0 {
			// TODO
			//service.Port = serviceContainer.Ports[0].Host
		}

		service.State = domain.ServiceStateRunning
	}

	return nil
}
