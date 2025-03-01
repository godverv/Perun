package domain

type Resource struct {
	Name        string
	Image       string
	ServiceName string
	State       serviceState
}

type ResourceDiff struct {
	StoppedResources []Resource
	NewResources     []Resource
}
