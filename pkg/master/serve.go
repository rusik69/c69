package master

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// Serve starts the server.
func Serve() {
	r := gin.New()
	r.GET("/api/v1/vm/:id", GetVMHandler)
	r.POST("/api/v1/vm/create", CreateVMHandler)
	r.DELETE("/api/v1/vm/:id", DeleteVMHandler)
	r.GET("/api/v1/vm/list", ListVMHandler)
	r.GET("/api/v1/vm/start/:id", StartVMHandler)
	r.GET("/api/v1/vm/stop/:id", StopVMHandler)
	r.GET("/api/v1/container/:id", GetContainerHandler)
	r.POST("/api/v1/container/create", CreateContainerHandler)
	r.DELETE("/api/v1/container/:id", DeleteContainerHandler)
	r.GET("/api/v1/container/list", ListContainerHandler)
	r.GET("/api/v1/container/start/:id", StartContainerHandler)
	r.GET("/api/v1/container/stop/:id", StopContainerHandler)
	r.POST("/api/v1/node/add", AddNodeHandler)
	r.GET("/api/v1/node/list", ListNodesHandler)
	r.GET("/api/v1/node/:id", GetNodeHandler)
	r.DELETE("/api/v1/node/:id", DeleteNodeHandler)
	logrus.Println("Master is listening on port " + string(types.MasterEnvInstance.ListenPort))
	r.Run(":" + types.MasterEnvInstance.ListenPort)
}
