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
	r.POST("/api/v1/vms", CreateVMHandler)
	r.GET("/api/v1/vm/:id", VMInfoHandler)
	r.DELETE("/api/v1/vm/:id", DeleteVMHandler)
	r.GET("/api/v1/vms", ListVMHandler)
	r.GET("api/v1/vm/start/:id", StartVMHandler)
	r.GET("api/v1/vm/stop/:id", StopVMHandler)
	r.GET("/api/v1/container/:id", GetContainerHandler)
	r.POST("/api/v1/container/create", CreateContainerHandler)
	r.DELETE("/api/v1/container/:id", DeleteContainerHandler)
	r.GET("/api/v1/containers", ListContainersHandler)
	r.POST("/api/v1/files", UploadFileHandler)
	r.DELETE("/api/v1/file/:id", FileDeleteHandler)
	r.GET("/api/v1/files", FileListHandler)
	r.GET("/api/v1/file/:id", FileGetHandler)
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
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempVM.Name == "" || tempVM.Image == "" || tempVM.Flavor == "" {
		logrus.Error("name, image or flavor is empty")
		c.JSON(400, gin.H{"error": "name, image or flavor is empty"})
		return
	}
	vm, err := CreateVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vm)
}

// DeleteHandler handles the delete request.
func DeleteVMHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tempVM := types.VM{ID: intID}
	err = DeleteVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// ListHandler handles the list request.
func ListVMHandler(c *gin.Context) {
	vms, err := ListVMs()
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vms)
}

// VMInfoHandler handles the get request.
func VMInfoHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tempVM := types.VM{ID: idInt}
	vm, err := GetVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vm)
}

// StopVMHandler handles the stop vm request.
func StopVMHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tempVM := types.VM{ID: idInt}
	err = StopVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// StartVMHandler handles the start vm request.
func StartVMHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tempVM := types.VM{ID: idInt}
	err = StartVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// CreateContainerHandler handles the create container request.
func CreateContainerHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempContainer types.Container
	if err := c.ShouldBindJSON(&tempContainer); err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempContainer.Name == "" || tempContainer.Image == "" {
		logrus.Error("name or image is empty")
		c.JSON(400, gin.H{"error": "name or image is empty"})
		return
	}
	container, err := CreateContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, container)
}

// DeleteContainerHandler handles the delete container request.
func DeleteContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	tempContainer := types.Container{ID: id}
	err := DeleteContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// GetContainerHandler handles the get container request.
func GetContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	tempContainer := types.Container{ID: id}
	container, err := GetContainer(tempContainer)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, container)
}

// ListContainersHandler handles the list container request.
func ListContainersHandler(c *gin.Context) {
	containers, err := ListContainers()
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, containers)
}

// UploadFileHandler handles the create file request.
func UploadFileHandler(c *gin.Context) {
	fileName := c.Query("name")
	if fileName == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}

}
