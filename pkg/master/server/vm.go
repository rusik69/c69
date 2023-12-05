package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/master/etcd"
	"github.com/rusik69/govnocloud/pkg/node/vm"
)

// CreateVMHandler handles the create vm request.
func CreateVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempVM vm.VM
	if err := c.ShouldBindJSON(&tempVM); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempVM.Name == "" || tempVM.Image == "" || tempVM.Flavor == "" {
		c.JSON(400, gin.H{"error": "name, image or flavor is empty"})
		return
	}
	vmInfoString, err := etcd.Get("/vms/" + tempVM.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	if vmInfoString != "" {
		c.JSON(400, gin.H{"error": "vm with this id already exists"})
	}
	tempVM.Committed = true
	tempVmstring, err := json.Marshal(tempVM)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	err = etcd.Put("/vms/"+tempVM.Name, string(tempVmstring))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	return
}
