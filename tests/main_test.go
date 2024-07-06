package tests

import (
	"context"
	"os"
	"testing"

	dockerCompose "github.com/harrim91/docker-compose-go/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Red-Sock/Perun/pkg/perun_api"
)

type TestEnv struct {
	perunApi perun_api.PerunAPIClient

	clean func()
}

var env TestEnv

func TestMain(m *testing.M) {
	var code int
	defer func() {
		os.Exit(code)
	}()

	//initCompose()
	//defer env.clean()

	initApi()

	code = m.Run()
}

func initCompose() {
	composeOpts := &dockerCompose.GlobalOptions{
		Files: []string{"./int_test.docker-compose.yaml"},
	}
	compose := dockerCompose.New(composeOpts)

	buildC, err := compose.Build(nil, os.Stdout)
	if err != nil {
		logrus.Fatal(err, "error building")
	}
	err = <-buildC
	if err != nil {
		logrus.Fatal(err, "error during build")
	}

	upOptions := &dockerCompose.UpOptions{
		Detach: true,
	}

	upC, err := compose.Up(upOptions, os.Stdout)
	if err != nil {
		logrus.Fatal(err, "error starting containers")
	}

	err = <-upC
	if err != nil {
		logrus.Fatal(err, "error during start")
	}

	env.clean = func() {
		dC, err := compose.Down(nil, os.Stdout)
		if err != nil {
			logrus.Fatal(err, "error cleaning up")
		}

		err = <-dC
		if err != nil {
			logrus.Fatal(err, "error during clean")
		}
	}
}

func initApi() {
	cl, err := grpc.NewClient("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}
	env.perunApi = perun_api.NewPerunAPIClient(cl)

	ctx := context.Background()

	securityDisabled := true

	connectVelezReq := &perun_api.ConnectVelez_Request{
		Node: &perun_api.Node{
			Name:             "int_test_velez",
			Addr:             "0.0.0.0:53890",
			SecurityDisabled: &securityDisabled,
		},
	}
	_, err = env.perunApi.ConnectVelez(ctx, connectVelezReq)
	if err != nil {
		logrus.Fatal(err)
	}
}
