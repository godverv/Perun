package run_service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/Red-Sock/Perun/internal/storage"
)

var ErrServiceAlreadyExists = errors.New("service already exists")

type Init struct {
	resourcesData storage.Resources
}

func NewInitStep(resources storage.Resources) *Init {
	return &Init{
		resourcesData: resources,
	}
}

func (i *Init) Do(ctx context.Context, r *RunServiceReq) error {
	service, err := i.resourcesData.Get(ctx, r.ServiceName)
	if err != nil {
		return errors.Wrap(err, "error getting service instance")
	}

	if service != nil {
		return errors.Wrap(ErrServiceAlreadyExists, "")
	}

	return nil
}
