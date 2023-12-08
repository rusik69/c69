package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/node/stats"
)

// StatsHandler handles the stats request.
func StatsHandler(c *gin.Context) {
	nodeStats, err := stats.Get()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nodeStats)
}
