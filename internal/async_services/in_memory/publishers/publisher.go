package publishers

import (
	"context"
	"sync"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/Red-Sock/Perun/internal/async_services"
)

var PublisherClosedErr = errors.New("publisher has been closed")

type Publisher[E any] struct {
	eventChan   chan E
	stopChan    chan struct{}
	subscribers []async_services.EventHandler[E]

	m sync.Mutex

	closed bool
}

func New[E any]() *Publisher[E] {
	p := &Publisher[E]{
		eventChan: make(chan E),
		stopChan:  make(chan struct{}),
	}
	p.start()

	return p
}

func (s *Publisher[E]) Subscribe(handler async_services.EventHandler[E]) {
	s.m.Lock()
	s.subscribers = append(s.subscribers, handler)
	s.m.Unlock()
}

func (s *Publisher[E]) Dispatch(event E) error {
	s.m.Lock()
	defer s.m.Unlock()

	if s.closed {
		return PublisherClosedErr
	}
	s.eventChan <- event

	return nil
}

func (s *Publisher[E]) Stop() {
	s.m.Lock()
	defer s.m.Unlock()
	if s.closed {
		return
	}

	s.closed = true

	close(s.stopChan)
}

func (s *Publisher[E]) start() {
	s.m.Lock()
	defer s.m.Unlock()
	if s.closed {
		return
	}

	go func() {
		for {
			select {
			case v := <-s.eventChan:
				s.consume(v)
			case _ = <-s.stopChan:
				return
			}
		}
	}()
}

func (s *Publisher[E]) consume(v E) {
	s.m.Lock()
	defer s.m.Unlock()
	errg, gctx := errgroup.WithContext(context.Background())
	for _, sub := range s.subscribers {
		errg.Go(func() error {
			return sub.Consume(gctx, v)
		})
	}

	err := errg.Wait()
	if err != nil {
		logrus.Error("error consuming event", err.Error())
	}
}
