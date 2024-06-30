package docker

import (
	errors "github.com/Red-Sock/trace-errors"
	docker "github.com/fsouza/go-dockerclient"
)

type Docker struct {
	client *docker.Client
}

func Init() (doc Docker, err error) {
	doc.client, err = docker.NewClientFromEnv()
	if err != nil {
		return doc, errors.Wrap(err, "error creating docker client")
	}

	return doc, nil
}
