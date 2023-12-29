package node

import (
	"context"
	"strconv"

	dockertypes "github.com/docker/docker/api/types"
	dockercontainer "github.com/docker/docker/api/types/container"
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
	dockerContainer := dockercontainer.Config{
		Image:  c.Image,
		Labels: map[string]string{"Name": c.Name},
	}
	resp, err := DockerConnection.ContainerCreate(ctx, &dockerContainer, nil, nil, nil, c.Name)
	if err != nil {
		return types.Container{}, err
	}
	IDstring, err := strconv.Atoi(resp.ID)
	if err != nil {
		return types.Container{}, err
	}
	c.ID = IDstring
	return c, nil
}

// DeleteContainer deletes a container.
func DeleteContainer(c types.Container) error {
	ctx := context.Background()
	idString := strconv.Itoa(c.ID)
	err := DockerConnection.ContainerRemove(ctx, idString, dockertypes.ContainerRemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

// StartContainer starts a container.
func StartContainer(c types.Container) error {
	ctx := context.Background()
	idString := strconv.Itoa(c.ID)
	err := DockerConnection.ContainerStart(ctx, idString, dockertypes.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

// StopContainer stops a container.
func StopContainer(c types.Container) error {
	ctx := context.Background()
	idString := strconv.Itoa(c.ID)
	err := DockerConnection.ContainerStop(ctx, idString, dockercontainer.StopOptions{})
	if err != nil {
		return err
	}
	return nil
}

// GetContainer gets a container.
func GetContainer(c types.Container) (types.Container, error) {
	ctx := context.Background()
	idString := strconv.Itoa(c.ID)
	container, err := DockerConnection.ContainerInspect(ctx, idString)
	if err != nil {
		return types.Container{}, err
	}
	c.Image = container.Config.Image
	c.Name = container.Name
	return c, nil
}

// ListContainers lists containers.
func ListContainers() ([]types.Container, error) {
	ctx := context.Background()
	containers, err := DockerConnection.ContainerList(ctx, dockertypes.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	var cs []types.Container
	for _, container := range containers {
		idInt, err := strconv.Atoi(container.ID)
		if err != nil {
			return nil, err
		}
		c := types.Container{
			ID:   idInt,
			Name: container.Labels["Name"],
		}
		cs = append(cs, c)
	}
	return cs, nil
}
