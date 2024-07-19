package domain

import (
	"github.com/godverv/Velez/pkg/velez_api"
)

type Service struct {
	Name     string
	Image    string
	State    serviceState
	Replicas int
}

type ServiceConstructor struct {
	Constructor *velez_api.CreateSmerd_Request
}
