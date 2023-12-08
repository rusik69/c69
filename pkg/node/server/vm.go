package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/node/vm"
)

// CreateHandler handles the create request.
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
	err := tempVM.Create()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// DeleteHandler handles the delete request.
func DeleteVMHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tempVM := vm.VM{ID: intID}
	err = tempVM.Delete()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// ListHandler handles the list request.
func ListVMHandler(c *gin.Context) {
	vms, err := vm.List()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vms)
}

// VMInfoHandler handles the get request.
func VMInfoHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tempVM := vm.VM{ID: idInt}
	err = tempVM.Get()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tempVM)
}
