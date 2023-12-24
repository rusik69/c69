package node

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
)

// Serve serves the node.
func Serve() {
	r := gin.New()
	r.POST("/api/v1/vm/create", CreateVMHandler)
	r.GET("/api/v1/vm/:id", VMInfoHandler)
	r.DELETE("/api/v1/vm/delete/:id", DeleteVMHandler)
	r.GET("/api/v1/vm/list", ListVMHandler)
	r.GET("/api/v1/node/stats", StatsHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	logrus.Println("Node is listening on port " + string(types.NodeEnvInstance.Port))
	r.Run(":" + string(types.NodeEnvInstance.Port))
}

// StatsHandler handles the stats request.
func StatsHandler(c *gin.Context) {
	nodeStats, err := GetStats()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nodeStats)
}

// Get gets the stats.
func GetStats() (types.Stats, error) {
	numCPUs, err := cpu.Counts(true)
	if err != nil {
		return types.Stats{}, err
	}
	mem, err := mem.VirtualMemory()
	if err != nil {
		return types.Stats{}, err
	}
	disk, err := disk.Usage("/")
	if err != nil {
		return types.Stats{}, err
	}
	return types.Stats{
		CPUs: numCPUs,
		MEM:  int64(mem.Total),
		DISK: int64(disk.Total),
	}, nil
}

// CreateHandler handles the create request.
func CreateVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempVM types.VM
	if err := c.ShouldBindJSON(&tempVM); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempVM.Name == "" || tempVM.Image == "" || tempVM.Flavor == "" {
		c.JSON(400, gin.H{"error": "name, image or flavor is empty"})
		return
	}
	err := CreateVM(tempVM)
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
	tempVM := types.VM{ID: intID}
	err = DeleteVM(tempVM)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// ListHandler handles the list request.
func ListVMHandler(c *gin.Context) {
	vms, err := ListVMs()
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
	tempVM := types.VM{ID: idInt}
	err = GetVM(tempVM)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tempVM)
}