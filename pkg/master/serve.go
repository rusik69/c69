package master

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// Serve starts the server.
func Serve() {
	r := gin.New()
	r.GET("/api/v1/vm/:name", GetVMHandler)
	r.POST("/api/v1/vms", CreateVMHandler)
	r.DELETE("/api/v1/vm/:id", DeleteVMHandler)
	r.GET("/api/v1/vms", ListVMHandler)
	r.GET("/api/v1/vm/start/:id", StartVMHandler)
	r.GET("/api/v1/vm/stop/:id", StopVMHandler)
	r.GET("/api/v1/container/:id", GetContainerHandler)
	r.POST("/api/v1/containers", CreateContainerHandler)
	r.DELETE("/api/v1/container/:id", DeleteContainerHandler)
	r.GET("/api/v1/containers", ListContainerHandler)
	r.GET("/api/v1/container/start/:id", StartContainerHandler)
	r.GET("/api/v1/container/stop/:id", StopContainerHandler)
	r.POST("/api/v1/nodes", AddNodeHandler)
	r.GET("/api/v1/nodes", ListNodesHandler)
	r.GET("/api/v1/node/:id", GetNodeHandler)
	r.DELETE("/api/v1/node/:id", DeleteNodeHandler)
	r.POST("/api/v1/files", PostFileHandler)
	r.DELETE("/api/v1/file/:name", DeleteFileHandler)
	r.GET("/api/v1/files", ListFilesHandler)
	r.GET("/api/v1/file/:name", GetFileHandler)
	logrus.Println("Master is listening on port " + string(types.MasterEnvInstance.ListenPort))
	r.Run(":" + types.MasterEnvInstance.ListenPort)
}
