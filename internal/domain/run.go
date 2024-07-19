package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type CreateServiceRequest struct {
	Constructor       *velez_api.CreateSmerd_Request
	ReplicationFactor uint32
}
