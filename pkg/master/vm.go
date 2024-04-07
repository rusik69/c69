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
	vmFlavorName := tempVM.Flavor
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
	nodes, err := GetNodes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	vmFlavor := types.VMFlavors[vmFlavorName]
	usedNode := types.Node{}
	for _, node := range nodes {
		if uint64(node.MilliCPUSTotal-node.MilliCPUSUsed) < vmFlavor.MilliCPUs ||
			(node.MemoryTotal-node.MemoryUsed) < vmFlavor.RAM ||
			(node.DiskTotal-node.DiskUsed) < vmFlavor.Disk {
			continue
		}
		createdVM, err := client.CreateVM(node.Host, node.Port, tempVM.Name,
			tempVM.Image, tempVM.Flavor)
		if err != nil {
			logrus.Error("can't create vm on node", node.Name)
			logrus.Error(err.Error())
			continue
		}
		newVM.ID = createdVM.ID
		newVM.Node = node.Name
		newVM.Name = tempVM.Name
		newVM.Image = tempVM.Image
		newVM.Flavor = tempVM.Flavor
		newVM.VNCURL = tempVM.VNCURL
		newVM.NodeHostname = tempVM.NodeHostname
		newVM.TailscaleID = createdVM.TailscaleID
		newVM.TailscaleIP = createdVM.TailscaleIP
		newVM.KubeConfig = createdVM.KubeConfig
		created = true
		usedNode = node
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
	usedNode.MilliCPUSUsed += vmFlavor.MilliCPUs
	usedNode.MemoryUsed += vmFlavor.RAM * 1024 * 1024
	usedNode.DiskUsed += vmFlavor.Disk * 1024 * 1024 * 1024
	nodeString, err := json.Marshal(usedNode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = ETCDPut("/nodes/"+usedNode.Name, string(nodeString))
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
	logrus.Printf("Deleting VM %s\n", vmInfoString)
	var vmInfo types.VM
	err = json.Unmarshal([]byte(vmInfoString), &vmInfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	deleted := false
	nodes, err := GetNodes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	foundNode := types.Node{}
	for _, node := range nodes {
		if node.Name == vmInfo.Node {
			err = client.DeleteVM(node.Host, node.Port, vmInfo.Name)
			if err != nil {
				logrus.Error(err.Error())
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			deleted = true
			foundNode = node
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
	foundNode.MilliCPUSUsed -= types.VMFlavors[vmInfo.Flavor].MilliCPUs
	foundNode.MemoryUsed -= types.VMFlavors[vmInfo.Flavor].RAM * 1024 * 1024
	foundNode.DiskUsed -= types.VMFlavors[vmInfo.Flavor].Disk * 1024 * 1024
	nodeString, err := json.Marshal(foundNode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	err = ETCDPut("/nodes/"+foundNode.Name, string(nodeString))
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
	vmsList, err := ETCDList("/vms/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	vmsMap := map[string]types.VM{}
	for _, vmName := range vmsList {
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
	var res []types.VM
	nodes, err := ETCDList("/nodes/")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	for _, nodeName := range nodes {
		logrus.Println("Node", nodeName)
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
		nodeVMS, err := client.ListVMs(node.Host, node.Port)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			logrus.Error("ListVMs error")
			logrus.Error(err.Error())
			return
		}
		for _, vm := range nodeVMS {
			if vmsMap[vm.Name].Name == vm.Name {
				var tempVM types.VM
				tempVM.ID = vm.ID
				tempVM.Name = vm.Name
				tempVM.IP = vm.IP
				tempVM.Node = nodeName
				tempVM.NodeHostname = node.Host
				tempVM.NodePort = node.Port
				tempVM.State = vm.State
				tempVM.Image = vmsMap[vm.Name].Image
				tempVM.Flavor = vmsMap[vm.Name].Flavor
				tempVM.Volumes = vmsMap[vm.Name].Volumes
				tempVM.VNCURL = vmsMap[vm.Name].VNCURL
				tempVM.TailscaleID = vmsMap[vm.Name].TailscaleID
				tempVM.TailscaleIP = vmsMap[vm.Name].TailscaleIP
				res = append(res, tempVM)
			}
		}
	}
	logrus.Println(res)
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
	nodes, err := GetNodes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	for _, node := range nodes {
		if node.Host == vmInfo.Node {
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
	nodes, err := GetNodes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		logrus.Error(err.Error())
		return
	}
	for _, node := range nodes {
		if node.Host == vmInfo.Node {
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
