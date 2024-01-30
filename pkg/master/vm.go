package master

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// CreateVMHandler handles the create vm request.
func CreateVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempVM types.VM
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
	logrus.Println("Creating VM", tempVM)
	vmInfoString, err := ETCDGet("/vms/" + tempVM.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString != "" {
		c.JSON(400, gin.H{"error": "vm with this name already exists"})
		logrus.Error("vm with this name already exists")
		return
	}
	created := false
	var newVM types.VM
	for _, node := range types.MasterEnvInstance.Nodes {
		newVMID, err := client.CreateVM(node.Host, node.Port, tempVM.Name, tempVM.Image, tempVM.Flavor)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		newVM.ID = newVMID
		newVM.Host = node.Name
		newVM.Name = tempVM.Name
		newVM.Image = tempVM.Image
		newVM.Flavor = tempVM.Flavor
		created = true
		break
	}
	if !created {
		c.JSON(500, gin.H{"error": "vm was not created"})
		logrus.Error("vm was not created")
		return
	}
	newVM.Committed = true
	newVmstring, err := json.Marshal(newVM)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = ETCDPut("/vms/"+newVM.Name, string(newVmstring))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, newVM)
}

// DeleteVMHandler handles the delete request.
func DeleteVMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		logrus.Error("name is empty")
		return
	}
	logrus.Printf("Deleting VM %s\n", name)
	vmInfoString, err := ETCDGet("/vms/" + name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString == "" {
		c.JSON(400, gin.H{"error": "vm with this name does not exist"})
		logrus.Error("vm with this name does not exist")
		return
	}
	var vmInfo types.VM
	err = json.Unmarshal([]byte(vmInfoString), &vmInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	deleted := false
	for _, node := range types.MasterEnvInstance.Nodes {
		if node.Host == vmInfo.Host {
			err = client.DeleteVM(node.Host, node.Port, vmInfo.Name)
			if err != nil {
				logrus.Error(err.Error())
				c.JSON(500, gin.H{"error": err.Error()})
				break
			}
			deleted = true
		}
	}
	if !deleted {
		c.JSON(500, gin.H{"error": "vm was not deleted"})
		logrus.Error("vm was not deleted")
		return
	}
	err = ETCDDelete("/vms/" + name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

// GetVMHandler handles the get vm info request.
func GetVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		logrus.Error("name is empty")
		return
	}
	logrus.Printf("Getting VM %s\n", name)
	vmInfoString, err := ETCDGet("/vms/" + name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString == "" {
		c.JSON(400, gin.H{"error": "vm with this name does not exist"})
		logrus.Error("vm with this name does not exist")
		return
	}
	var vmInfo types.VM
	err = json.Unmarshal([]byte(vmInfoString), &vmInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	c.JSON(200, vmInfo)
}

// ListVMHandler handles the list vm request.
func ListVMHandler(c *gin.Context) {
	logrus.Println("Listing VMs")
	nodesList, err := ETCDList("/nodes/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	vms, err := ETCDList("/vms/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	vmsMap := map[string]types.VM{}
	for _, vmName := range vms {
		vmString, err := ETCDGet(vmName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		var vm types.VM
		err = json.Unmarshal([]byte(vmString), &vm)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		vmsMap[vm.Name] = vm
	}
	logrus.Println("VMs map:", vmsMap)
	res := []types.VM{}
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
		nodeVMs, err := client.ListVMs(node.Host, node.Port)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error(err.Error())
			return
		}
		for _, nodeVM := range nodeVMs {
			var resVM types.VM
			resVM.Name = nodeVM.Name
			resVM.Image = vmsMap[nodeVM.Name].Image
			resVM.Flavor = vmsMap[nodeVM.Name].Flavor
			resVM.ID = nodeVM.ID
			resVM.Host = vmsMap[nodeVM.Name].Host
			resVM.State = nodeVM.State
			res = append(res, resVM)
		}
	}
	logrus.Println("VMs:", res)
	c.JSON(200, res)
}

// StartVMHandler handles the start vm request.
func StartVMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		logrus.Error("name is empty")
		return
	}
	logrus.Printf("Starting VM %s\n", name)
	vmInfoString, err := ETCDGet("/vms/" + name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString == "" {
		c.JSON(400, gin.H{"error": "vm with this name does not exist"})
		logrus.Error("vm with this name does not exist")
		return
	}
	var vmInfo types.VM
	err = json.Unmarshal([]byte(vmInfoString), &vmInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	for _, node := range types.MasterEnvInstance.Nodes {
		if node.Host == vmInfo.Host {
			err = client.StartVM(node.Host, node.Port, vmInfo.Name)
			if err != nil {
				logrus.Error(err.Error())
				c.JSON(500, gin.H{"error": err.Error()})
				break
			}
		}
	}
	c.JSON(200, gin.H{"status": "ok"})
}

// StopVMHandler handles the stop vm request.
func StopVMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		logrus.Error("name is empty")
		return
	}
	logrus.Printf("Stopping VM %s\n", name)
	vmInfoString, err := ETCDGet("/vms/" + name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	if vmInfoString == "" {
		c.JSON(400, gin.H{"error": "vm with this name does not exist"})
		logrus.Error("vm with this name does not exist")
		return
	}
	var vmInfo types.VM
	err = json.Unmarshal([]byte(vmInfoString), &vmInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	for _, node := range types.MasterEnvInstance.Nodes {
		if node.Host == vmInfo.Host {
			err = client.StopVM(node.Host, node.Port, vmInfo.Name)
			if err != nil {
				logrus.Error(err.Error())
				c.JSON(500, gin.H{"error": err.Error()})
				break
			}
		}
	}
	c.JSON(200, gin.H{"status": "ok"})
}
