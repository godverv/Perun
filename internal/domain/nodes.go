package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type PickNodeReq struct {
	ReplicationFactor uint32
}

type Node struct {
	Name string
	Conn velez_api.VelezAPIClient
}
