package run_service

import (
	"context"
	"time"

	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/storage"
)

var defaultServiceUpWaitTime = time.Second * 10

type PostRunStep struct {
	resourcesStorage storage.Resources
}

func (p *PostRunStep) Do(ctx context.Context, r *RunServiceReq) error {
	time.Sleep(max(r.Config.StartupDuration, defaultServiceUpWaitTime))

	listReq := &velez_api.ListSmerds_Request{}

	listReq.Name = &r.ServiceName
	for _, node := range r.Nodes {
		list, err := node.Conn.ListSmerds(ctx, listReq)
		if err != nil {
			// TODO: write error to db
			continue
		}

		var serviceContainer *velez_api.Smerd
		for _, smerd := range list.GetSmerds() {
			if smerd.GetName() == r.ServiceName {
				serviceContainer = smerd
				break
			}
		}
		if serviceContainer == nil {
			// TODO: write error to db
			continue
		}

		if serviceContainer.Status != velez_api.Smerd_running {
			// TODO: write error to db
		}
	}

	return nil
}
