package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type Dependencies struct {
	Smerds []Dependency
}

type Dependency struct {
	Name     string
	SmerdReq *velez_api.CreateSmerd_Request
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
	Port         uint32
}
