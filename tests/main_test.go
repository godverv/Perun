package tests

import (
	"os"
	"testing"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/Perun/internal/utils/test_utils/docker"
)

type TestEnv struct {
	dockerClient *docker.Docker
}

var env TestEnv

func TestMain(m *testing.M) {
	var err error
	env, err = initTestEnv()
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
	var code int
	defer os.Exit(code)

	defer env.dockerClient.TearDown()

	code = m.Run()
}

func initTestEnv() (env TestEnv, err error) {
	env.dockerClient, err = docker.Init()
	if err != nil {
		return env, errors.Wrap(err, "error initializing docker client")
	}

	return env, nil
}
