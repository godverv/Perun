package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/Red-Sock/Perun/pkg/perun_api"
)

type PublishServiceSuite struct {
	suite.Suite

	ctx context.Context
}

func (p *PublishServiceSuite) SetupSuite() {
	p.ctx = context.Background()
}

func (p *PublishServiceSuite) Test_OK() {
	runServiceReq := &perun_api.RunService_Request{
		ImageName:         "redsockdev/red-cart:v0.0.13",
		ServiceName:       "red-cart",
		ReplicationFactor: 1,
	}

	resp, err := env.perunApi.RunService(p.ctx, runServiceReq)
	require.NoError(p.T(), err)
	_ = resp
}

func Test_PublishServiceSuite(t *testing.T) {
	suite.Run(t, new(PublishServiceSuite))
}
