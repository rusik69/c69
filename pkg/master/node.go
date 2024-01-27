package master

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// AddNodeHandler handles the add node request.
func AddNodeHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempNode types.Node
	if err := c.ShouldBindJSON(&tempNode); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if tempNode.Name == "" || tempNode.Host == "" || tempNode.Port == "" {
		c.JSON(400, gin.H{"error": "name, host and port are required"})
		logrus.Error("name, host and port are required")
		return
	}
	tempNodeString, err := json.Marshal(tempNode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = ETCDPut("/nodes/"+tempNode.Name, string(tempNodeString))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}

// ListNodesHandler handles the list nodes request.
func ListNodesHandler(c *gin.Context) {
	nodes, err := ETCDList("/nodes/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, nodes)
}

// GetNodeHandler handles the get node request.
func GetNodeHandler(c *gin.Context) {
	nodeName := c.Param("name")
	if nodeName == "" {
		c.JSON(400, gin.H{"error": "node name is required"})
		logrus.Error("node name is required")
		return
	}
	nodeString, err := ETCDGet("/nodes/" + nodeName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if nodeString == "" {
		c.JSON(404, gin.H{"error": "node not found"})
		logrus.Error("node not found")
		return
	}
	c.JSON(200, gin.H{"node": nodeString})
}

// DeleteNodeHandler handles the delete node request.
func DeleteNodeHandler(c *gin.Context) {
	nodeName := c.Param("name")
	if nodeName == "" {
		c.JSON(400, gin.H{"error": "node name is required"})
		logrus.Error("node name is required")
		return
	}
	err := ETCDDelete("/nodes/" + nodeName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}

// AddEnvNodesToETCD adds nodes from env to etcd.
func AddEnvNodesToETCD() error {
	nodesList, err := ETCDList("/nodes/")
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	nodesMap := map[string]bool{}
	for _, node := range nodesList {
		nodesMap[node] = true
	}
	for _, node := range types.MasterEnvInstance.Nodes {
		if nodesMap[node.Name] {
			logrus.Printf("Node %s already exists in etcd\n", node.Name)
			continue
		}
		nodeString, err := json.Marshal(node)
		if err != nil {
			logrus.Error(err.Error())
			return err
		}
		err = ETCDPut("/nodes/"+node.Name, string(nodeString))
		if err != nil {
			logrus.Error(err.Error())
			return err
		}
	}
	return nil
}
