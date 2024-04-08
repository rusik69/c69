package node

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateLLMHandler handles the create LLM request.
func CreateLLMHandler(c *gin.Context) {
	var llm types.LLM
	err := c.BindJSON(&llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = CreateLLM(llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, llm)
}

// GetLLMHandler gets a llm cluster
func GetLLMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	logrus.Println("Getting LLM", name)
	llm, err := GetLLM(name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, llm)
}

// DeleteLLMHandler handles the delete LLM request.
func DeleteLLMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	err := DeleteLLM(name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "llm deleted"})
}

// ListLLMHandler handles the list LLM request.
func ListLLMHandler(c *gin.Context) {
	llms, err := ListLLM()
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, llms)
}

// StartLLMHandler handles the start LLM request.
func StartLLMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	err := StartLLM(name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "llm started"})
}

// StopLLMHandler handles the stop LLM request.
func StopLLMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	err := StopLLM(name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "llm stopped"})
}
