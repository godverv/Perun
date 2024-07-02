package cron

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Cron struct {
	stop func()
}

func New(duration time.Duration, do func()) *Cron {
	ticker := time.NewTicker(duration)
	do()
	go func() {
		for range ticker.C {
			do()
		}
	}()

	c := &Cron{
		stop: ticker.Stop,
	}

	return c
}

func (c *Cron) Stop() {
	if c.stop != nil {
		c.stop()
	} else {
		logrus.Error("no stop function defined")
	}
}

func (c *Cron) start() {

}
