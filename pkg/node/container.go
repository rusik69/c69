package node

import (
	"context"

	dockertypes "github.com/docker/docker/api/types"
	dockercontainer "github.com/docker/docker/api/types/container"
	dockerclient "github.com/docker/docker/client"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

var DockerConnection *dockerclient.Client

// ContainerConnect connects to the container daemon.
func ContainerConnect() (*dockerclient.Client, error) {
	cli, err := dockerclient.NewClientWithOpts(dockerclient.FromEnv)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// CreateContainer creates a container.
func CreateContainer(c types.Container) (types.Container, error) {
	logrus.Println("Creating container", c.Name, "with image", c.Image)
	ctx := context.Background()
	pullOptions := dockertypes.ImagePullOptions{}
	_, err := DockerConnection.ImagePull(ctx, c.Image, pullOptions)
	if err != nil {
		return types.Container{}, err
	}
	dockerContainer := dockercontainer.Config{
		Image:  c.Image,
		Labels: map[string]string{"Name": c.Name},
	}
	resp, err := DockerConnection.ContainerCreate(ctx, &dockerContainer, nil, nil, nil, c.Name)
	if err != nil {
		return types.Container{}, err
	}
	c.ID = resp.ID
	return c, nil
}

// DeleteContainer deletes a container.
func DeleteContainer(c types.Container) error {
	logrus.Println("Deleting container", c.Name)
	ctx := context.Background()
	err := DockerConnection.ContainerRemove(ctx, c.ID, dockertypes.ContainerRemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

// StartContainer starts a container.
func StartContainer(c types.Container) error {
	logrus.Println("Starting container", c.Name)
	ctx := context.Background()
	err := DockerConnection.ContainerStart(ctx, c.ID, dockertypes.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

// StopContainer stops a container.
func StopContainer(c types.Container) error {
	logrus.Println("Stopping container", c.Name)
	ctx := context.Background()
	err := DockerConnection.ContainerStop(ctx, c.ID, dockercontainer.StopOptions{})
	if err != nil {
		return err
	}
	return nil
}

// GetContainer gets a container.
func GetContainer(c types.Container) (types.Container, error) {
	logrus.Println("Getting container", c.Name)
	ctx := context.Background()
	container, err := DockerConnection.ContainerInspect(ctx, c.ID)
	if err != nil {
		return types.Container{}, err
	}
	c.Image = container.Config.Image
	c.Name = container.Name
	return c, nil
}

// ListContainers lists containers.
func ListContainers() ([]types.Container, error) {
	logrus.Println("Listing containers")
	ctx := context.Background()
	containers, err := DockerConnection.ContainerList(ctx, dockertypes.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	var cs []types.Container
	for _, container := range containers {
		c := types.Container{
			ID:   container.ID,
			Name: container.Labels["Name"],
		}
		cs = append(cs, c)
	}
	return cs, nil
}
