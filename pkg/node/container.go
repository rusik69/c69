package node

import (
	"context"
	"errors"
	"io"

	dockercontainer "github.com/docker/docker/api/types/container"
	dockerimage "github.com/docker/docker/api/types/image"
	dockerclient "github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

var DockerConnection *dockerclient.Client

// CreateContainerHandler handles the create container request.
func CreateContainerHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempContainer types.Container
	if err := c.ShouldBindJSON(&tempContainer); err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempContainer.Name == "" || tempContainer.Image == "" {
		logrus.Error("name or image is empty")
		c.JSON(400, gin.H{"error": "name or image is empty"})
		return
	}
	logrus.Println("Creating container", tempContainer.Name, tempContainer.Image)
	container, err := CreateContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// check if container name ends with -llm
	if len(tempContainer.Name) > 4 && tempContainer.Name[len(tempContainer.Name)-4:] == "-llm" {
		err = waitForLLM(container.IP)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, container)
}

// DeleteContainerHandler handles the delete container request.
func DeleteContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	tempContainer := types.Container{ID: id}
	logrus.Println("Deleting container", tempContainer.ID)
	_ = StopContainer(tempContainer)
	err := DeleteContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// GetContainerHandler handles the get container request.
func GetContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	tempContainer := types.Container{ID: id}
	container, err := GetContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, container)
}

// StartContainerHandler handles the start container request.
func StartContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	tempContainer := types.Container{ID: id}
	logrus.Println("Starting container", tempContainer.ID)
	err := StartContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// StopContainerHandler handles the stop container request.
func StopContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	tempContainer := types.Container{ID: id}
	logrus.Println("Stopping container", tempContainer.ID)
	err := StopContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// ListContainersHandler handles the list container request.
func ListContainersHandler(c *gin.Context) {
	containers, err := ListContainers()
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, containers)
}

// ContainerConnect connects to the container daemon.
func ContainerConnect() (*dockerclient.Client, error) {
	logrus.Println("Connecting to docker")
	cli, err := dockerclient.NewClientWithOpts(dockerclient.FromEnv)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// CreateContainer creates a container.
func CreateContainer(c types.Container) (types.Container, error) {
	ctx := context.Background()
	reader, err := DockerConnection.ImagePull(ctx, c.Image, dockerimage.PullOptions{})
	if err != nil {
		return types.Container{}, err
	}
	defer reader.Close()
	_, err = io.Copy(io.Discard, reader)
	if err != nil {
		return types.Container{}, err
	}
	dockerContainer := dockercontainer.Config{
		Image:  c.Image,
		Labels: map[string]string{"Name": c.Name},
	}
	memLimit := types.ContainerFlavors[c.Flavor].Mem * 1024 * 1024
	cpuShares := types.ContainerFlavors[c.Flavor].MilliCPUs
	hostConfig := dockercontainer.HostConfig{
		Resources: dockercontainer.Resources{
			Memory:    int64(memLimit),
			CPUShares: int64(cpuShares),
		},
		NetworkMode: "br0",
	}
	resp, err := DockerConnection.ContainerCreate(ctx, &dockerContainer, &hostConfig, nil, nil, c.Name)
	if err != nil {
		return types.Container{}, err
	}
	err = StartContainer(types.Container{ID: resp.ID})
	if err != nil {
		return types.Container{}, err
	}
	runningContainer, err := FindContainerByName(c.Name)
	if err != nil {
		return types.Container{}, err
	}
	c.ID = resp.ID
	c.IP = runningContainer.IP
	return c, nil
}

// DeleteContainer deletes a container.
func DeleteContainer(c types.Container) error {
	ctx := context.Background()
	err := DockerConnection.ContainerRemove(ctx, c.ID, dockercontainer.RemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

// StartContainer starts a container.
func StartContainer(c types.Container) error {
	ctx := context.Background()
	err := DockerConnection.ContainerStart(ctx, c.ID, dockercontainer.StartOptions{})
	if err != nil {
		return err
	}
	return nil
}

// StopContainer stops a container.
func StopContainer(c types.Container) error {
	ctx := context.Background()
	err := DockerConnection.ContainerStop(ctx, c.ID, dockercontainer.StopOptions{})
	if err != nil {
		return err
	}
	return nil
}

// GetContainer gets a container.
func GetContainer(c types.Container) (types.Container, error) {
	ctx := context.Background()
	container, err := DockerConnection.ContainerInspect(ctx, c.ID)
	if err != nil {
		return types.Container{}, err
	}
	c.Image = container.Config.Image
	c.Name = container.Name
	c.IP = container.NetworkSettings.Networks["bridge"].IPAddress
	return c, nil
}

// FindContainerByName finds a container by name.
func FindContainerByName(name string) (types.Container, error) {
	ctx := context.Background()
	containers, err := DockerConnection.ContainerList(ctx, dockercontainer.ListOptions{})
	if err != nil {
		return types.Container{}, err
	}
	for _, container := range containers {
		logrus.Printf("Checking container %+v", container)
		if container.Labels["Name"] == name {
			logrus.Println("Found container", container)
			c := types.Container{
				ID:   container.ID,
				Name: container.Labels["Name"],
				IP:   container.NetworkSettings.Networks["bridge"].IPAddress,
			}
			return c, nil
		}
	}
	return types.Container{}, errors.New("container not found")
}

// ListContainers lists containers.
func ListContainers() ([]types.Container, error) {
	ctx := context.Background()
	containers, err := DockerConnection.ContainerList(ctx, dockercontainer.ListOptions{})
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
