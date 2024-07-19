package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type PickNodesReq struct {
	NodesCount uint32
}

type Node struct {
	Name string
	Conn velez_api.VelezAPIClient
}
