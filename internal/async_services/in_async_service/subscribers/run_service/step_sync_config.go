package run_service

import (
	"context"
	"fmt"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type SyncConfigStep struct {
	matreshka matreshka_api.MatreshkaBeAPIClient
}

func NewSyncConfigStep(matreshka matreshka_api.MatreshkaBeAPIClient) *SyncConfigStep {
	return &SyncConfigStep{
		matreshka: matreshka,
	}
}

func (s *SyncConfigStep) Do(ctx context.Context, r *RunServiceReq) error {
	patchConfigReq := &matreshka_api.PatchConfig_Request{}
	patchConfigReq.ServiceName = r.ServiceName

	for _, resource := range r.Dependencies.Resources {
		for _, configChange := range resource.ConfigChanges {

			valueStr := fmt.Sprint(configChange.Value)
			newConfigValue := &matreshka_api.Node{
				Name:  resource.Name,
				Value: &valueStr,
			}

			patchConfigReq.Changes = append(patchConfigReq.Changes, newConfigValue)
		}
	}
	_, err := s.matreshka.PatchConfig(ctx, patchConfigReq)
	if err != nil {
		return errors.Wrap(err, "error patching config")
	}

	return nil
}
