package run_service

import (
	"context"
	"sync"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/storage"
)

type RunServiceReq struct {
	ServiceName string
	ImageName   string

	ReplicationFactor uint32

	Nodes []domain.Node

	Config matreshka.AppConfig
}

type Step interface {
	Do(ctx context.Context, r *RunServiceReq) error
}

type ServiceRunner struct {
	steps []Step

	q chan RunServiceReq

	maxParallel int
}

func New(
	services service.Services,
	data storage.Data,
) *ServiceRunner {
	return &ServiceRunner{
		steps: []Step{
			NewInitStep(data.Resources()),
			NewGetNodesStep(services.Nodes()),
			NewPreSyncConfigStep(),
			NewSyncDependenciesStep(services.Resources(), data.Resources()),
			NewRunServiceStep(data.Resources()),
		},

		maxParallel: 2,
	}
}

func (s *ServiceRunner) Run(ctx context.Context) chan RunServiceReq {

	q := make(chan RunServiceReq)

	go func() {
		wg := sync.WaitGroup{}
		maxChan := make(chan struct{}, s.maxParallel)
		defer close(maxChan)

		for {
			select {
			case req := <-q:
				wg.Add(1)
				maxChan <- struct{}{}

				go func() {
					defer wg.Done()
					defer func() {
						_ = <-maxChan
					}()
					s.run(ctx, &req)
				}()

			case <-ctx.Done():

				wg.Wait()
				err := ctx.Err()
				if err != nil {
					logrus.Error(errors.Wrap(err, "error from ctx"))
				}
			}
		}
	}()
	return q
}

func (s *ServiceRunner) run(ctx context.Context, req *RunServiceReq) {
	for _, step := range s.steps {
		err := step.Do(ctx, req)
		if err != nil {
			logrus.Error(errors.Wrap(err, "error running step"))
			return
		}
	}
}
