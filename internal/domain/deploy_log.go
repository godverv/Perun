package domain

import (
	"time"
)

type DeployLog struct {
	Id        int
	Name      string
	State     serviceState
	Reason    string
	CreatedAt time.Time
}
