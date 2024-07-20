package deploy_service

import (
	"context"
	"strings"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type DeployService struct {
	resourceSrv service.ResourceService
	deployLog   storage.DeployLogs
}

func New(data storage.Data, srv service.Services) *DeployService {
	return &DeployService{
		resourceSrv: srv.Resources(),

		deployLog: data.DeployLogs(),
	}
}

func (d *DeployService) Consume(ctx context.Context, req domain.DeployServiceReq) error {
	diff, err := d.resourceSrv.GetDiffForService(ctx, req.ServiceName)
	if err != nil {
		return errors.Wrap(err, "error getting diff for service's resources")
	}

	if len(diff.StoppedResources) != 0 || len(diff.NewResources) != 0 {
		return d.diffInResources(ctx, req, diff)
	}

	return nil
}

func (d *DeployService) diffInResources(ctx context.Context, req domain.DeployServiceReq, diff domain.ResourceDiff) error {
	log := domain.DeployLog{
		Name:  req.ServiceName,
		State: domain.ErrorDuringDeploy,
	}
	sb := strings.Builder{}

	sb.WriteString("Resource difference:\n")

	if len(diff.NewResources) != 0 {
		sb.WriteString("New resources:\n")
		for _, r := range diff.NewResources {
			sb.WriteString("- ")
			sb.WriteString(r.Name)
			sb.WriteByte('\n')
		}
	}

	if len(diff.StoppedResources) != 0 {
		sb.WriteString("Stopped resources:\n")
		for _, r := range diff.StoppedResources {
			sb.WriteString("- ")
			sb.WriteString(r.Name)
			sb.WriteByte('\n')
		}
	}

	log.Reason = sb.String()

	err := d.deployLog.Add(ctx, log)
	if err != nil {
		return errors.Wrap(err, "error saving log")
	}

	return nil
}
