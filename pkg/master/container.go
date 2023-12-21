package master

import "github.com/gin-gonic/gin"

// CreateContainerHandler handles the create container request.
func CreateContainerHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
}

// DeleteContainerHandler handles the delete container request.
func DeleteContainerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
}

// ListContainerHandler handles the list container request.
func ListContainerHandler(c *gin.Context) {
}

// GetContainerHandler handles the get container request.
func GetContainerHandler(c *gin.Context) {
}
