package domain

import (
	"github.com/Red-Sock/evon"
	"github.com/godverv/Velez/pkg/velez_api"
)

type Dependencies struct {
	Volumes   []Volume
	Resources []DataSource
}

type DataSource struct {
	Name          string
	Constructor   *velez_api.CreateSmerd_Request
	ConfigChanges []evon.Node
}

type Volume struct {
	Binding *velez_api.VolumeBindings
}
