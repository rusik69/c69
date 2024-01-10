package node

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// PostFileHandler handles the create file request.
func PostFileHandler(c *gin.Context) {
	fileName := c.Query("name")
	if fileName == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	file := c.Request.Body
	defer file.Close()
	err := SaveFile(fileName, file)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

// GetFileHandler handles the download file request.
func GetFileHandler(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	fileName := filepath.Join(types.NodeEnvInstance.FilesDir, name)
	c.File(fileName)
}

// DeleteFileHandler handles the delete file request.
func DeleteFileHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	err := DeleteFile(name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// ListFilesHandler handles the list files request.
func ListFilesHandler(c *gin.Context) {
	files, err := ListFiles()
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, files)
}

// SaveFile saves a file.
func SaveFile(name string, file io.Reader) error {
	fileName := filepath.Join(types.NodeEnvInstance.FilesDir, name)
	logrus.Println("SaveFile: ", fileName)
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile deletes a file.
func DeleteFile(name string) error {
	fileName := filepath.Join(types.NodeEnvInstance.FilesDir, name)
	logrus.Println("DeleteFile: ", fileName)
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

// ListFiles lists files.
func ListFiles() ([]string, error) {
	files := []string{}
	err := filepath.Walk(types.NodeEnvInstance.FilesDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, info.Name())
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
