package perun

import (
	"context"

	"github.com/Red-Sock/Perun/pkg/perun_api"
)

func (impl *Implementation) Version(_ context.Context, _ *perun_api.Version_Request) (*perun_api.Version_Response, error) {
	return &perun_api.Version_Response{
		Version: impl.version,
	}, nil
}
