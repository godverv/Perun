package run_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"
	"github.com/godverv/matreshka"
)

type SyncConfigStep struct{}

func NewPreSyncConfigStep() *SyncConfigStep {
	return &SyncConfigStep{}
}

func (s *SyncConfigStep) Do(ctx context.Context, r *RunServiceReq) error {
	fetchReq := &velez_api.FetchConfig_Request{
		ServiceName: r.ServiceName,
		ImageName:   r.ImageName,
	}

	if len(r.Nodes) == 0 {
		return errors.Wrap(ErrNoNodes, "no nodes provided to this step")
	}

	fetchResp, err := r.Nodes[0].Conn.FetchConfig(ctx, fetchReq)
	if err != nil {
		return errors.Wrap(err, "error fetching config")
	}

	r.Config = matreshka.NewEmptyConfig()
	err = r.Config.Unmarshal(fetchResp.Config)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling fetched config")
	}

	return nil
}
