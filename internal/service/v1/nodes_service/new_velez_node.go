package nodes_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/data/velez"
	"github.com/Red-Sock/Perun/internal/domain"
)

const defaultVelezKeyPath = "/tmp/velez/private.key"

func (n *nodesService) NewVelezConnection(ctx context.Context, in domain.VelezConnection) (err error) {
	if in.Node.CustomVelezKeyPath == "" {
		in.Node.CustomVelezKeyPath = defaultVelezKeyPath
	}

	vs, err := velez.NewVelezService(in)
	if err != nil {
		return errors.Wrap(err, "error connecting to velez node")
	}

	_, err = vs.Version(ctx, &velez_api.Version_Request{})
	if err != nil {
		return errors.Wrap(err, "error getting velez version")
	}

	err = n.nodesStore.SaveConnection(ctx, in)
	if err != nil {
		return errors.Wrap(err, "error saving node info to stores")
	}

	return nil
}
