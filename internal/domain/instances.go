package domain

type Instance struct {
	Name     string
	NodeName string
	Port     int
	State    serviceState
	Image    string
}
