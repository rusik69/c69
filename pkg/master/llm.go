package master

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// CreateLLMHandler handles the create LLM request.
func CreateLLMHandler(c *gin.Context) {
	var tempLLM types.LLM
	if err := c.ShouldBindJSON(&tempLLM); err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempLLM.Name == "" || tempLLM.Model == "" {
		logrus.Error("name or flavor is empty")
		c.JSON(400, gin.H{"error": "name or flavor is empty"})
		return
	}
	logrus.Println("Creating LLM", tempLLM)
	llmInfoString, err := ETCDGet("/llm/" + tempLLM.Name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if llmInfoString != "" {
		logrus.Error("llm with this name already exists")
		c.JSON(400, gin.H{"error": "llm with this name already exists"})
		return
	}
	image := types.LLMModels[tempLLM.Model].Image
	containerFlavor := types.LLMModels[tempLLM.Model].ContainerFlavor
	ctrID, err := client.CreateContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort,
		tempLLM.Name+"-llm", image, containerFlavor)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	llm := types.LLM{
		Name:  tempLLM.Name,
		Model: tempLLM.Model,
		Container: types.Container{
			ID:     ctrID,
			Name:   tempLLM.Name + "-llm",
			Image:  image,
			Flavor: containerFlavor,
		},
	}
	llmString, err := json.Marshal(llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = ETCDPut("/llm/"+tempLLM.Name, string(llmString))
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, llm)
}

// GetLLMHandler handles the get LLM request.
func GetLLMHandler(c *gin.Context) {
	name := c.Param("name")
	llmInfoString, err := ETCDGet("/llm/" + name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logrus.Println("Getting LLM", name)
	if llmInfoString == "" {
		logrus.Error("llm with this name does not exist")
		c.JSON(400, gin.H{"error": "llm with this name does not exist"})
		return
	}
	var llm types.LLM
	err = json.Unmarshal([]byte(llmInfoString), &llm)
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
	logrus.Println("Deleting LLM", name)
	llmString, err := ETCDGet("/llm/" + name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if llmString == "" {
		logrus.Error("llm with this name does not exist")
		c.JSON(400, gin.H{"error": "llm with this name does not exist"})
		return
	}
	var llm types.LLM
	err = json.Unmarshal([]byte(llmString), &llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = DeleteLLM(llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = ETCDDelete("/llm/" + name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "llm deleted"})
}

// DeleteLLM deletes a llm cluster.
func DeleteLLM(llm types.LLM) error {
	client.StopContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, llm.Container.Name)
	err := client.DeleteContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, llm.Container.Name)
	if err != nil {
		return err
	}
	return nil
}

// ListLLMsHandler handles the list LLMs request.
func ListLLMsHandler(c *gin.Context) {
	logrus.Println("Listing LLMs")
	llmNames, err := ListLLMs()
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var llms []types.LLM
	for _, name := range llmNames {
		llmString, err := ETCDGet(name)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var llm types.LLM
		err = json.Unmarshal([]byte(llmString), &llm)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		llms = append(llms, llm)
	}
	c.JSON(200, llms)
}

// ListLLMs lists all llm clusters.
func ListLLMs() ([]string, error) {
	llmList, err := ETCDList("/llm")
	if err != nil {
		return nil, err
	}
	return llmList, nil
}

// StartLLMHandler handles the start LLM request.
func StartLLMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	logrus.Println("Starting LLM", name)
	llmString, err := ETCDGet("/llm/" + name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if llmString == "" {
		logrus.Error("llm with this name does not exist")
		c.JSON(400, gin.H{"error": "llm with this name does not exist"})
		return
	}
	var llm types.LLM
	err = json.Unmarshal([]byte(llmString), &llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = StartLLM(llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "llm started"})
}

// StartLLM starts a llm cluster.
func StartLLM(llm types.LLM) error {
	err := client.StartContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, llm.Container.Name)
	if err != nil {
		return err
	}
	return nil
}

// StopLLM stops a llm cluster.
func StopLLM(llm types.LLM) error {
	err := client.StopContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, llm.Container.Name)
	if err != nil {
		return err
	}
	return nil
}

// StopLLMHandler handles the stop LLM request.
func StopLLMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	logrus.Println("Stopping LLM", name)
	llmString, err := ETCDGet("/llm/" + name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if llmString == "" {
		logrus.Error("llm with this name does not exist")
		c.JSON(400, gin.H{"error": "llm with this name does not exist"})
		return
	}
	var llm types.LLM
	err = json.Unmarshal([]byte(llmString), &llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = StopLLM(llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "llm stopped"})
}

// GenerateLLMHandler handles the generate LLM request.
func GenerateLLMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	logrus.Println("Generating LLM", name)
	llmString, err := ETCDGet("/llm/" + name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if llmString == "" {
		logrus.Error("llm with this name does not exist")
		c.JSON(400, gin.H{"error": "llm with this name does not exist"})
		return
	}
	var llm types.LLM
	err = json.Unmarshal([]byte(llmString), &llm)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logrus.Printf("llm: %+v\n", llm)
	// read request body to string msg
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	input := string(bodyBytes)
	msg, err := GenerateLLM(llm, input)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, msg)
}

// GenerateLLM gets a response from llm.
func GenerateLLM(llm types.LLM, input string) (string, error) {
	nodeName := llm.Container.Host
	node, err := GetNode(nodeName)
	if err != nil {
		return "", err
	}
	host := node.Host
	port := node.Port
	containerName := llm.Container.Name
	url := "http://" + host + ":" + port + "/api/v1/llmgenerate/" + containerName
	req, err := http.NewRequest("POST", url, strings.NewReader(input))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(string(bodyText))
	}
	return string(bodyText), nil
}
