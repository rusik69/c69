package master

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// PostFileHandler handles the post file request.
func PostFileHandler(c *gin.Context) {
	var tempFile types.File
	body := c.Request.Body
	defer body.Close()
	if err := c.ShouldBindJSON(&tempFile); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if tempFile.Name == "" || tempFile.Size == 0 {
		c.JSON(400, gin.H{"error": "name or size is empty"})
		logrus.Error("name or size is empty")
		return
	}
	logrus.Println("Creating file", tempFile)
	fileInfoString, err := ETCDGet("/files/" + tempFile.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if fileInfoString == "" {
		node, err := chooseNodeForFile(tempFile)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		tempFile.Host = node.Host
	}
}

func chooseNodeForFile(file types.File) (types.Node, error) {
	found := false
	var foundNode types.Node
	for _, node := range types.MasterEnvInstance.Nodes {
		nodeStats, err := client.GetNodeStats(node.Host, node.Port)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		if nodeStats.FreeDISK > file.Size {
			found = true
			foundNode = node
			break
		}
	}
	if found {
		return foundNode, nil
	} else {
		return types.Node{}, errors.New("can't find node for file")
	}
}
