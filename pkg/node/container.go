package node

import (
	"context"

	dockertypes "github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
	"github.com/rusik69/govnocloud/pkg/types"
)

var DockerConnection *dockerclient.Client

// ContainerConnect connects to the container daemon.
func ContainerConnect() {
	cli, err := dockerclient.NewClientWithOpts(dockerclient.FromEnv)
	if err != nil {
		panic(err)
	}
	DockerConnection = cli
}

// CreateContainer creates a container.
func CreateContainer(c types.Container) (types.Container, error) {
	ctx := context.Background()
	dockerContainer := dockertypes.Container{
		Image: c.Image,
		Names : []string{c.Name},
	}
	resp, err := DockerConnection.ContainerCreate(ctx, dockerContainer, nil, nil, name)
	if err != nil {
		return err
	}
	_, 
	return nil
}
