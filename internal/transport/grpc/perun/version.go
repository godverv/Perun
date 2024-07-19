package perun

import (
	"context"

	"github.com/Red-Sock/Perun/pkg/perun_api"
)

func (s *Impl) Version(_ context.Context, _ *perun_api.Version_Request) (*perun_api.Version_Response, error) {
	return &perun_api.Version_Response{
		Version: s.version,
	}, nil
}
