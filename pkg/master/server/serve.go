package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/node/env"
	"github.com/sirupsen/logrus"
)

// Serve starts the server.
func Serve() {
	r := gin.New()
	r.GET("/api/v1/vm/:id", GetVMHandler)
	r.POST("/api/v1/vm/create", CreateVMHandler)
	r.DELETE("/api/v1/vm/:id", DeleteVMHandler)
	r.GET("/api/v1/vm/list", ListVMHandler)
	r.POST("/api/v1/node/add", AddNodeHandler)
	r.GET("/api/v1/node/list", ListNodesHandler)
	r.GET("/api/v1/node/:id", GetNodeHandler)
	r.DELETE("/api/v1/node/:id", DeleteNodeHandler)
	logrus.Println("Master is listening on port " + string(env.NodeEnvInstance.Port))
}
