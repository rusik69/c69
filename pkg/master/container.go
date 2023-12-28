package master

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// CreateContainerHandler handles the create container request.
func CreateContainerHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempContainer types.Container
	if err := c.ShouldBindJSON(&tempContainer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if tempContainer.Name == "" || tempContainer.Image == "" {
		c.JSON(400, gin.H{"error": "name or image is empty"})
		logrus.Error("name or image is empty")
		return
	}
	logrus.Println("Creating container", tempContainer)
	containerInfoString, err := ETCDGet("/containers/" + tempContainer.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if containerInfoString != "" {
		c.JSON(400, gin.H{"error": "container with this id already exists"})
		logrus.Error("container with this id already exists")
		return
	}
	newContainerID := 0
	created := false
	var newContainer types.Container
	for _, node := range types.MasterEnvInstance.Nodes {
		
}

// DeleteContainerHandler handles the delete container request.
func DeleteContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	logrus.Printf("Deleting container with id %s\n", id)
	var tempContainer types.Container
	tempContainer.ID = id

	err := DeleteContainer(tempContainer)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

}

// ListContainerHandler handles the list container request.
func ListContainerHandler(c *gin.Context) {

}

// GetContainerHandler handles the get container request.
func GetContainerHandler(c *gin.Context) {
}
