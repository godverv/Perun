package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type Dependencies struct {
	Smerds  []*velez_api.CreateSmerd_Request
	Volumes []*velez_api.VolumeBindings
}

type resourceState int

const (
	ResourceStateInvalid resourceState = iota
	ResourceStateCreated
	ResourceStateRunning
	ResourceStateError
)

type Resource struct {
	ResourceName string
	NodeName     string
	State        resourceState
}

type UpdateState struct {
	ResourceName string
	State        resourceState
}
