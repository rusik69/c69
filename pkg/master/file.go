package master

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// UploadFileHandler handles the post file request.
func UploadFileHandler(c *gin.Context) {
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
	logrus.Println("Uploading file", tempFile)
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
		tempFile.NodeHost = node.Host
		tempFile.NodePort = node.Port
		tempFile.NodeName = node.Name
		fileInfoBytes, err := json.Marshal(tempFile)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		err = ETCDPut("/files/"+tempFile.Name, string(fileInfoBytes))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		c.JSON(200, node)
	} else {
		c.JSON(400, gin.H{"error": "file already exists"})
		logrus.Error("file already exists")
		return
	}
}

// CommitFileHandler handles the commit file request.
func CommitFileHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		logrus.Error("name is empty")
		return
	}
	fileInfoString, err := ETCDGet("/files/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	logrus.Println("Committing file", name)
	if fileInfoString == "" {
		c.JSON(400, gin.H{"error": "file " + name + " not found"})
		logrus.Error("file " + name + " not found")
		return
	}
	var fileInfo types.File
	err = json.Unmarshal([]byte(fileInfoString), &fileInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	fileInfo.Committed = true
	newFileInfoString, err := json.Marshal(fileInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = ETCDPut("/files/"+name, string(newFileInfoString))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

// DeleteFileHandler handles the delete file request.
func DeleteFileHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		logrus.Error("name is empty")
		return
	}
	logrus.Println("Deleting file", name)
	fileInfoString, err := ETCDGet("/files/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if fileInfoString == "" {
		c.JSON(400, gin.H{"error": "file not found"})
		logrus.Error("file not found")
		return
	}
	var fileInfo types.File
	err = json.Unmarshal([]byte(fileInfoString), &fileInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	fileNode, err := GetNode(fileInfo.NodeName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	logrus.Println("Deleting file", name, "from node", fileNode)
	err = client.DeleteFile(fileNode.Host, fileNode.Port, name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = ETCDDelete("/files/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

// ListFilesHandler handles the list files request.
func ListFilesHandler(c *gin.Context) {
	filesList, err := ETCDList("/files/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	var files []types.File
	for _, file := range filesList {
		fileInfoString, err := ETCDGet(file)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		var fileInfo types.File
		err = json.Unmarshal([]byte(fileInfoString), &fileInfo)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		files = append(files, fileInfo)
	}
	c.JSON(200, files)
}

// GetFileHandler handles the get file request.
func GetFileHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		logrus.Error("name is empty")
		return
	}
	fileInfoString, err := ETCDGet("/files/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if fileInfoString == "" {
		c.JSON(400, gin.H{"error": "file not found"})
		logrus.Error("file not found")
		return
	}
	var fileInfo types.File
	err = json.Unmarshal([]byte(fileInfoString), &fileInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, fileInfo)
}

// chooseNodeForFile chooses the node for file.
func chooseNodeForFile(file types.File) (types.Node, error) {
	found := false
	var foundNode types.Node
	nodes, err := GetNodes()
	if err != nil {
		return types.Node{}, err
	}
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
	for _, node := range nodes {
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
