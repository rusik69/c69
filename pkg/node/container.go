package node

import (
	dockerclient "github.com/docker/docker/client"
)

var DockerConnection *dockerclient.Client

// ContainerConnect connects to the container daemon.
func ContainerConnect() {
	cli, err := dockerclient.NewEnvClient()
	if err != nil {
		panic(err)
	}
	DockerConnection = cli
}

// CreateContainer creates a container.
func CreateContainer(name string) error {
	_, err := DockerConnection.ContainerCreate(nil, nil, nil, nil, name)
	if err != nil {
		return err
	}
	return nil
}