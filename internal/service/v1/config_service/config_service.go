package config_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
)

type ConfigService struct {
	ns service.NodesService
}

func New(ns service.NodesService) *ConfigService {
	return &ConfigService{
		ns: ns,
	}
}

func (c *ConfigService) FetchForService(ctx context.Context, srv domain.Service) (matreshka.AppConfig, error) {
	pickNodesReq := domain.PickNodesReq{NodesCount: 1}

	cfg := matreshka.NewEmptyConfig()

	node, err := c.ns.PickNodes(ctx, pickNodesReq)
	if err != nil {
		return cfg, errors.Wrap(err, "error picking nodes")
	}

	fetchReq := &velez_api.FetchConfig_Request{
		ServiceName: srv.Name,
		ImageName:   srv.Image,
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
