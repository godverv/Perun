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
	SmerdReq      *velez_api.CreateSmerd_Request
	ConfigChanges []evon.Node
}
type Volume struct {
	Binding *velez_api.VolumeBindings
}

type resourceState int

const (
	ResourceStateInvalid resourceState = iota
	ResourceStateCreated
	ResourceStateWaitingForResources
	ResourceStateRunning
	ResourceStateStopped
	ResourceStateError
)

type Resource struct {
	ResourceName string
	NodeName     string
	State        resourceState
	Port         uint32
}

type ListResources struct {
	Name string
}
