package sync_service_info

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
)

func (r *RefreshServiceConfig) syncConfig(ctx context.Context, req domain.SyncServiceInfo) (matreshka.AppConfig, error) {
	pickNodesReq := domain.PickNodesReq{NodesCount: 1}

	cfg := matreshka.NewEmptyConfig()

	node, err := r.nodesService.PickNodes(ctx, pickNodesReq)
	if err != nil {
		return cfg, errors.Wrap(err, "error picking nodes")
	}

	fetchReq := &velez_api.FetchConfig_Request{
		ServiceName: req.Service.Name,
		ImageName:   req.Service.Image,
	}

	fetchResp, err := node[0].Conn.FetchConfig(ctx, fetchReq)
	if err != nil {
		return cfg, errors.Wrap(err, "error fetching config")
	}

	err = cfg.Unmarshal(fetchResp.Config)
	if err != nil {
		return cfg, errors.Wrap(err, "error unmarshalling fetched config")
	}

	return cfg, nil
}
