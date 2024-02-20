package master

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/client"
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
	failed := true
	count := 0
	for failed {
		if count == 10 {
			c.JSON(500, gin.H{"error": "node is not available"})
			logrus.Error("node is not available")
			return
		}
		req, err := http.Get("http://" + tempNode.Host + ":" + tempNode.Port + "/ping")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		if req.StatusCode != 200 {
			c.JSON(500, gin.H{"error": "node is not available"})
			logrus.Error("node is not available")
			return
		} else {
			failed = false
		}
		count++
		time.Sleep(1 * time.Second)
	}
	nodeStats, err := client.GetNodeStats(tempNode.Host, tempNode.Port)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	tempNode.MilliCPUSTotal = nodeStats.MilliCPUs
	tempNode.MemoryTotal = nodeStats.TotalMEM
	tempNode.DiskTotal = nodeStats.TotalDISK
	logrus.Println("Adding node", tempNode)
	tempNodeBody, err := json.Marshal(tempNode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = ETCDPut("/nodes/"+tempNode.Name, string(tempNodeBody))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}

// ListNodesHandler handles the list nodes request.
func ListNodesHandler(c *gin.Context) {
	logrus.Println("Listing nodes")
	nodesList, err := ETCDList("/nodes/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	res := map[string]types.NodeStats{}
	for _, nodeName := range nodesList {
		nodeString, err := ETCDGet(nodeName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		var node types.Node
		err = json.Unmarshal([]byte(nodeString), &node)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		nodeStats, err := client.GetNodeStats(node.Host, node.Port)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		res[nodeName] = nodeStats
	}
	c.JSON(200, res)
}

// GetNodeHandler handles the get node request.
func GetNodeHandler(c *gin.Context) {
	nodeName := c.Param("name")
	if nodeName == "" {
		c.JSON(400, gin.H{"error": "node name is required"})
		logrus.Error("node name is required")
		return
	}
	logrus.Println("Getting node", nodeName)
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
	var node types.Node
	err = json.Unmarshal([]byte(nodeString), &node)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, node)
}

// DeleteNodeHandler handles the delete node request.
func DeleteNodeHandler(c *gin.Context) {
	nodeName := c.Param("name")
	if nodeName == "" {
		c.JSON(400, gin.H{"error": "node name is required"})
		logrus.Error("node name is required")
		return
	}
	logrus.Println("Deleting node", nodeName)
	err := ETCDDelete("/nodes/" + nodeName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}
