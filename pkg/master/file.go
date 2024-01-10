package master

import (
	"github.com/gin-gonic/gin"
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
	if tempFile.Name == "" || tempFile.Size == "" {
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

	}
}

func chooseNodeForFile() (types.Node, error) {
	found := false
	var foundNode types.Node
	for _, node := range types.MasterEnvInstance.Nodes {

		if node.Capacity > 0 {
			found = true
			foundNode = node
			break
		}
	}
}
