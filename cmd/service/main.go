package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	grpcResource "github.com/Red-Sock/Perun/internal/clients/grpc"
	"github.com/Red-Sock/Perun/internal/clients/sqlite"
	"github.com/Red-Sock/Perun/internal/config"
	"github.com/Red-Sock/Perun/internal/data"
	"github.com/Red-Sock/Perun/internal/data/storage"
	"github.com/Red-Sock/Perun/internal/service"
	"github.com/Red-Sock/Perun/internal/service/v1"
	"github.com/Red-Sock/Perun/internal/transport/grpc"
	"github.com/Red-Sock/Perun/internal/transport/grpc/perun"
	"github.com/Red-Sock/Perun/internal/utils/closer"
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

	if cfg.GetAppInfo().StartupDuration == 0 {
		return errors.New("no startup duration in config")
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.GetAppInfo().StartupDuration)
	closer.Add(func() error { cancel(); return nil })

	matreshkaBeClient, err := grpcResource.NewMatreshkaBeAPIClient(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "error initializing matreshka client")
	}

	store, err := initStorage(cfg)
	if err != nil {
		return errors.Wrap(err, "error initializing storage")
	}

	srv, err := v1.NewService(ctx, store, matreshkaBeClient)
	if err != nil {
		return errors.Wrap(err, "error initializing node service")
	}

	err = runGrpcServer(ctx, cfg, srv)
	if err != nil {
		return errors.Wrap(err, "error running server")
	}

	waitingForTheEnd()

	logrus.Println("shutting down the app")

	return closer.Close()
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is a LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func initStorage(cfg config.Config) (data.Storage, error) {
	sqliteConn, err := cfg.GetDataSources().Sqlite(config.ResourceSqlite)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing sqlite client")
	}

	provider, err := sqlite.NewProvider(sqliteConn)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing sqlite storage")
	}

	s, err := storage.NewStorage(provider)
	if err != nil {
		return nil, errors.Wrap(err, "error creating storage")
	}
	return s, nil
}

func runGrpcServer(ctx context.Context, cfg config.Config, srv service.Services) error {
	grpcServerCfg, err := cfg.GetServers().GRPC(config.ServerGrpc)
	if err != nil {
		return errors.Wrap(err, "error getting grpc server config")
	}

	perunImp := perun.New(cfg, srv.Nodes(), srv.Runner())

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
