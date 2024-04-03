package node

import (
	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())
	r.POST("/api/v1/vms", CreateVMHandler)
	r.GET("/api/v1/vm/:id", GetVMHandler)
	r.DELETE("/api/v1/vm/:name", DeleteVMHandler)
	r.GET("/api/v1/vms", ListVMHandler)
	r.GET("api/v1/vmstart/:name", StartVMHandler)
	r.GET("api/v1/vmstop/:name", StopVMHandler)
	r.GET("/api/v1/container/:id", GetContainerHandler)
	r.POST("/api/v1/containers", CreateContainerHandler)
	r.DELETE("/api/v1/container/:id", DeleteContainerHandler)
	r.GET("/api/v1/containerstart/:id", StartContainerHandler)
	r.GET("/api/v1/containerstop/:id", StopContainerHandler)
	r.GET("/api/v1/containers", ListContainersHandler)
	r.POST("/api/v1/file/:name", PostFileHandler)
	r.DELETE("/api/v1/file/:name", DeleteFileHandler)
	r.GET("/api/v1/files", ListFilesHandler)
	r.GET("/api/v1/file/:name", GetFileHandler)
	r.GET("/api/v1/stats", StatsHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	logrus.Println("Node is listening on port " + string(types.NodeEnvInstance.ListenPort))
	err := r.Run(types.NodeEnvInstance.ListenHost + ":" + string(types.NodeEnvInstance.ListenPort))
	if err != nil {
		panic(err)
	}
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
func GetStats() (types.NodeStats, error) {
	numCPUs, err := cpu.Counts(true)
	if err != nil {
		return types.NodeStats{}, err
	}
	memory, err := mem.VirtualMemory()
	if err != nil {
		return types.NodeStats{}, err
	}
	usage, err := disk.Usage("/")
	if err != nil {
		return types.NodeStats{}, err
	}
	return types.NodeStats{
		MilliCPUs: uint64(1024 * numCPUs),
		FreeMEM:   uint64(memory.Free),
		TotalMEM:  uint64(memory.Total),
		FreeDISK:  uint64(usage.Free),
		TotalDISK: uint64(usage.Total),
	}, nil
}
