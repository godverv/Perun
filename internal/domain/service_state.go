package domain

const (
	ServiceStateInvalid serviceState = iota
	ServiceStateCreated
	ServiceStateStarting

	ServiceStateRunningOk
	ServiceStateRunningPartially
	ErrorDuringDeploy
	ErrorAlreadyDeployed
)

type serviceState int
