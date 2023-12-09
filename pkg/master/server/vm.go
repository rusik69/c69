package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	nodeClient "github.com/rusik69/govnocloud/pkg/client/node"
	"github.com/rusik69/govnocloud/pkg/master/env"
	"github.com/rusik69/govnocloud/pkg/master/etcd"
	"github.com/rusik69/govnocloud/pkg/node/vm"
	"github.com/sirupsen/logrus"
)

// CreateVMHandler handles the create vm request.
func CreateVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempVM vm.VM
	if err := c.ShouldBindJSON(&tempVM); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if tempVM.Name == "" || tempVM.Image == "" || tempVM.Flavor == "" {
		c.JSON(400, gin.H{"error": "name, image or flavor is empty"})
		logrus.Error("name, image or flavor is empty")
		return
	}
	vmInfoString, err := ETCDGet("/vms/" + tempVM.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString != "" {
		c.JSON(400, gin.H{"error": "vm with this id already exists"})
		logrus.Error("vm with this id already exists")
		return
	}
	newVMID := 0
	created := false
	var newVM vm.VM
	for _, node := range env.MasterEnvInstance.Nodes {
		newVMID, err = nodeClient.CreateVM(node.Host, node.Port, tempVM.Name, tempVM.Image, tempVM.Flavor)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		newVM.ID = newVMID
		newVM.Host = node.Host
		created = true
	}
	if !created {
		c.JSON(500, gin.H{"error": "vm was not created"})
		logrus.Error("vm was not created")
		return
	}
	newVM.Committed = true
	newVmstring, err := json.Marshal(newVM)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = etcd.Put("/vms/"+newVM.Name, string(newVmstring))
	if err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	return
}

// DeleteVMHandler handles the delete request.
func DeleteVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempVM vm.VM
	if err := c.ShouldBindJSON(&tempVM); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if tempVM.ID == 0 {
		c.JSON(400, gin.H{"error": "id is empty"})
		logrus.Error("id is empty")
		return
	}
	vmInfoString, err := ETCDGet("/vms/" + string(tempVM.ID))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString == "" {
		c.JSON(400, gin.H{"error": "vm with this id does not exist"})
		logrus.Error("vm with this id does not exist")
		return
	}
	var vmInfo vm.VM
	err = json.Unmarshal([]byte(vmInfoString), &vmInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	for _, node := range env.MasterEnvInstance.Nodes {
		if node.Host == vmInfo.Host {
			err = nodeClient.DeleteVM(tempVM.ID, node.Host, node.Port)
			if err != nil {
				logrus.Error(err.Error())
				c.JSON(500, gin.H{"error": err.Error()})
				break
			}
		}
	}
	err = etcd.Delete("/vms/" + string(tempVM.ID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
	return
}

// GetVMHandler handles the get vm info request.
func GetVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempVM vm.VM
	if err := c.ShouldBindJSON(&tempVM); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if tempVM.ID == 0 {
		c.JSON(400, gin.H{"error": "id is empty"})
		logrus.Error("id is empty")
		return
	}
	vmInfoString, err := ETCDGet("/vms/" + string(tempVM.ID))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString == "" {
		c.JSON(400, gin.H{"error": "vm with this id does not exist"})
		logrus.Error("vm with this id does not exist")
		return
	}
	var vmInfo vm.VM
	err = json.Unmarshal([]byte(vmInfoString), &vmInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, vmInfo)
	return
}

// ListVMHandler handles the list vm request.
func ListVMHandler(c *gin.Context) {
	vms, err := ETCDList("/vms/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, vms)
	return
}
