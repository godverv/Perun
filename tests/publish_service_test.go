package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PublishServiceSuite struct {
	suite.Suite
}

func (p *PublishServiceSuite) SetupSuite() {
	env.dockerClient.LaunchVelez(p.T())
}

func (p *PublishServiceSuite) Test_OK() {

}

func Test_PublishServiceSuite(t *testing.T) {
	suite.Run(t, new(PublishServiceSuite))
}
