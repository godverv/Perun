package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/Perun/internal/async_services"
	"github.com/Red-Sock/Perun/internal/async_services/in_async_service"
	grpcResource "github.com/Red-Sock/Perun/internal/clients/grpc"
	"github.com/Red-Sock/Perun/internal/clients/sqlite"
	"github.com/Red-Sock/Perun/internal/config"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/service/v1"
	"github.com/Red-Sock/Perun/internal/storage"
	"github.com/Red-Sock/Perun/internal/storage/data"
	"github.com/Red-Sock/Perun/internal/tasks/warm_up_cache"
	"github.com/Red-Sock/Perun/internal/transport/grpc"
	"github.com/Red-Sock/Perun/internal/transport/grpc/perun"
	"github.com/Red-Sock/Perun/internal/utils/closer"
	"github.com/Red-Sock/Perun/internal/utils/cron"
	//_transport_imports
)

func main() {
	err := start()
	if err != nil {
		logrus.Fatal(err)
	}
}

func start() error {
	logrus.Println("starting app")

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return errors.Wrap(err, "error reading config")
	}

	ctx, cancel := context.WithCancel(ctx)
	closer.Add(func() error { cancel(); return nil })

	matreshkaBeClient, err := grpcResource.NewMatreshkaBeAPIClient(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "error initializing matreshka client")
	}
	// TODO
	_ = matreshkaBeClient
	store, err := initStorage(cfg)
	if err != nil {
		return errors.Wrap(err, "error initializing storage")
	}

	srv := v1.NewService(store)

	cronCacheWarmUp(store)

	as := in_async_service.New(store, srv)
	closer.Add(as.Stop)

	err = runGrpcServer(ctx, cfg, srv, as)
	if err != nil {
		return errors.Wrap(err, "error running server")
	}

	waitingForTheEnd()

	logrus.Println("shutting down the app")

	return closer.Close()
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is an LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func initStorage(cfg config.Config) (storage.Data, error) {
	var provider *sql.DB
	{
		sqliteConn, err := cfg.GetDataSources().Sqlite(config.ResourceSqlite)
		if err != nil {
			return nil, errors.Wrap(err, "error initializing sqlite client")
		}

		provider, err = sqlite.NewProvider(sqliteConn)
		if err != nil {
			return nil, errors.Wrap(err, "error initializing sqlite storage")
		}
	}

	s, err := data.NewStorage(provider)
	if err != nil {
		return nil, errors.Wrap(err, "error creating storage")
	}

	return s, nil
}

func runGrpcServer(ctx context.Context, cfg config.Config, srv service.Services, queue async_services.AsyncService) error {
	grpcServerCfg, err := cfg.GetServers().GRPC(config.ServerGrpc)
	if err != nil {
		return errors.Wrap(err, "error getting grpc server config")
	}

	perunImp := perun.New(cfg, srv.Nodes(), queue)

	grpcApi, err := grpc.NewServer(grpcServerCfg, perunImp)
	if err != nil {
		return errors.Wrap(err, "error creating server")
	}

	err = grpcApi.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "error starting server")
	}

	closer.Add(func() error { return grpcApi.Stop(ctx) })

	return nil
}

func cronCacheWarmUp(d storage.Data) {
	cron.New(time.Minute, warm_up_cache.New(d).Do)
}
