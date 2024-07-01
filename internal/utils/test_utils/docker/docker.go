package docker

import (
	"context"
	"sync"
	"testing"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	docker "github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/stretchr/testify/require"
)

const (
	velezContainerName = "velez_test"
	velezImageName     = "velez:local"

	interTestLabel = "inter_test"
)

type Docker struct {
	client *docker.Client

	velez sync.Once
}

func Init() (doc *Docker, err error) {
	doc = &Docker{}
	doc.client, err = docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		return nil, errors.Wrap(err, "error creating docker client")
	}

	return doc, nil
}

func (d *Docker) LaunchVelez(t *testing.T) {
	d.velez.Do(func() {
		hostConfig := &container.HostConfig{
			PortBindings: nat.PortMap{
				"53890": []nat.PortBinding{
					{
						HostPort: "53890",
					},
				},
			},
			Mounts: []mount.Mount{
				{
					Type:   "bind",
					Source: "/var/run/docker.sock",
					Target: "/var/run/docker.sock",
				},
			},
		}
		contConf := &container.Config{
			Labels: map[string]string{
				interTestLabel: "true",
			},
			Hostname: velezContainerName,
			Image:    velezImageName,
			Env:      []string{"VELEZ_DISABLE_API_SECURITY=true"},
		}
		require.NoError(t, d.launchContainer(velezImageName, contConf, hostConfig))
	})
}

func (d *Docker) launchContainer(imageName string, contConf *container.Config, hostConfig *container.HostConfig) error {
	ctx := context.Background()

	c, err := d.client.ContainerInspect(ctx, velezContainerName)
	if err != nil {
		if !errdefs.IsNotFound(err) {
			return errors.Wrap(err, "error inspecting container")
		}
	}

	if c.ContainerJSONBase != nil {
		return nil
	}

	listFilter := filters.NewArgs()
	listFilter.Add("reference", imageName)

	imageListReq := image.ListOptions{
		All:     true,
		Filters: listFilter,
	}
	list, err := d.client.ImageList(ctx, imageListReq)
	if err != nil {
		return errors.Wrap(err, "error listing images")
	}

	if len(list) == 0 {
		_, err = d.client.ImagePull(ctx, imageName, image.PullOptions{})
		if err != nil {
			return errors.Wrap(err, "error pulling image")
		}
	}

	cont, err := d.client.ContainerCreate(ctx,
		contConf,
		hostConfig,
		&network.NetworkingConfig{},
		&v1.Platform{},
		contConf.Hostname,
	)
	if err != nil {
		return errors.Wrap(err, "error creating container")
	}

	err = d.client.ContainerStart(ctx, cont.ID, container.StartOptions{})
	if err != nil {
		return errors.Wrap(err, "error starting container")
	}

	return nil
}

func (d *Docker) TearDown() error {
	ctx := context.Background()

	listFilter := filters.NewArgs()
	listFilter.Add("label", interTestLabel+"=true")

	listReq := container.ListOptions{
		All:     true,
		Filters: listFilter,
	}
	list, err := d.client.ContainerList(ctx, listReq)
	if err != nil {
		return errors.Wrap(err, "error listing containers")
	}

	for _, testCont := range list {
		err = d.client.ContainerRemove(ctx, testCont.ID, container.RemoveOptions{
			RemoveVolumes: true,
			Force:         true,
		})
		if err != nil {
			return errors.Wrap(err, "error removing container")
		}
	}

	return nil
}
