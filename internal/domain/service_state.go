package domain

const (
	ServiceStateInvalid serviceState = iota
	ServiceStateCreated
	ServiceStateStarting

	ServiceStateRunningOk
	ServiceStateRunningPartially
	ServiceStateErrorDuringDeploy
	ServiceStateErrorAlreadyDeployed
	ServiceStateDeployDeleted
)

type serviceState int
