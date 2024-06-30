package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type RunServiceRequest struct {
	Constructor *velez_api.CreateSmerd_Request
	Nodes       []Node
}
