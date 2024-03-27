package master

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// CreateK8SHandler creates a k8s cluster
func CreateK8SHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempK8S types.K8S
	if err := c.ShouldBindJSON(&tempK8S); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempK8S.Name == "" || tempK8S.Flavor == "" {
		c.JSON(400, gin.H{"error": "name or flavor is empty"})
		return
	}
	vmFlavorName := tempK8S.Flavor
	logrus.Println("Creating K8S", tempK8S)
	k8sInfoString, err := ETCDGet("/k8s/" + tempK8S.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if k8sInfoString != "" {
		c.JSON(400, gin.H{"error": "k8s with this name already exists"})
		return
	}
	vm := types.VM{
		Name:   tempK8S.Name + "-k8s",
		Flavor: vmFlavorName,
		Image:  "k8s",
	}
	newVM, err := client.CreateVM(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort,
		vm.Name, vm.Image, vm.Flavor)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	tempK8S.VM = newVM
	tempK8SString, err := json.Marshal(tempK8S)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = ETCDPut("/k8s/"+tempK8S.Name, string(tempK8SString))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tempK8S)
}

// GetK8SHandler gets a k8s cluster
func GetK8SHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	logrus.Println("Getting K8S", name)
	k8sInfoString, err := ETCDGet("/k8s/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if k8sInfoString == "" {
		c.JSON(404, gin.H{"error": "k8s with this name does not exist"})
		return
	}
	var k8s types.K8S
	err = json.Unmarshal([]byte(k8sInfoString), &k8s)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, k8s)
}

// DeleteK8SHandler deletes a k8s cluster
func DeleteK8SHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	logrus.Println("Deleting K8S", name)
	k8sInfoString, err := ETCDGet("/k8s/" + name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if k8sInfoString == "" {
		c.JSON(400, gin.H{"error": "k8s with this name does not exist"})
		return
	}
	var k8s types.K8S
	err = json.Unmarshal([]byte(k8sInfoString), &k8s)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = client.DeleteVM(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, k8s.VM.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = ETCDDelete("/k8s/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "k8s deleted"})
}

// ListK8SHandler lists k8s clusters
func ListK8SHandler(c *gin.Context) {
	k8sList, err := ETCDList("/k8s")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res := []types.K8S{}
	for _, k8sName := range k8sList {
		k8sString, err := ETCDGet(k8sName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var k8s types.K8S
		err = json.Unmarshal([]byte(k8sString), &k8s)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		res = append(res, k8s)
	}
	c.JSON(200, res)
}

// StartK8SHandler starts a k8s cluster
func StartK8SHandler(c *gin.Context) {
	name := c.Param("name")
	k8sInfoString, err := ETCDGet("/k8s/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if k8sInfoString == "" {
		c.JSON(404, gin.H{"error": "k8s with this name does not exist"})
		return
	}
	var k8s types.K8S
	err = json.Unmarshal([]byte(k8sInfoString), &k8s)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = client.StartVM(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, k8s.VM.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "k8s started"})
}

// StopK8SHandler stops a k8s cluster
func StopK8SHandler(c *gin.Context) {
	name := c.Param("name")
	k8sInfoString, err := ETCDGet("/k8s/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if k8sInfoString == "" {
		c.JSON(404, gin.H{"error": "k8s with this name does not exist"})
		return
	}
	var k8s types.K8S
	err = json.Unmarshal([]byte(k8sInfoString), &k8s)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = client.StopVM(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, k8s.VM.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "k8s stopped"})
}
