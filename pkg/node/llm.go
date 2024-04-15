package node

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GenerateLLMHandler handles the generate LLM request.
func GenerateLLMHandler(c *gin.Context) {
	containerName := c.Param("name")
	if containerName == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	logrus.Println("Generating LLM response", containerName)
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	input := string(bodyBytes)
	ctr, err := FindContainerByName(containerName)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logrus.Println("llm ctr: %+v", ctr)
	url := "http://" + ctr.IP + "/generate"
	resp, err := http.Post(url, "text/plain", strings.NewReader(input))
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if resp.StatusCode != 200 {
		logrus.Error(string(bodyText))
		c.JSON(500, gin.H{"error": string(bodyText)})
		return
	}
	c.JSON(200, string(bodyText))
}
