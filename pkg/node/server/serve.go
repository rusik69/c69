package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/node/env"
	"github.com/sirupsen/logrus"
)

// Serve serves the node.
func Serve() {
	r := gin.New()
	r.POST("/api/v1/vm/create", CreateVMHandler)
	r.DELETE("/api/v1/vm/delete/:id", DeleteVMHandler)
	r.GET("/api/v1/vm/list", ListVMHandler)
	logrus.Println("Node is listening on port " + string(env.NodeEnvInstance.Port))
	r.Run(":" + string(env.NodeEnvInstance.Port))
}
