package domain

const (
	ServiceStateInvalid serviceState = iota
	ServiceStateCreated
	ServiceStateStarting

	ServiceStateRunningOk
	ServiceStateRunningPartially
	ErrorDuringDeploy
)

type serviceState int
