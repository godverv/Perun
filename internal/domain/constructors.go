package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type ResourceConstructor struct {
	Name      string
	BaseImage string

	DeployRequest *velez_api.CreateSmerd_Request
}
