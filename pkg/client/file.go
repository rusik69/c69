package client

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/rusik69/govnocloud/pkg/types"
)

// UploadFile uploads a file.
func UploadFile(masterHost, masterPort, name, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	url := "http://" + masterHost + ":" + masterPort + "/api/v1/files"
	var tempFile types.File
	tempFile.Name = name
	tempFile.Size = file.Stat().Size()
	tempFile.Timestamp = time.Now().Unix()
	tempFileBody, err := json.Marshal(tempFile)
	if err != nil {
		return err
	}
	if err := http.Post(url); err != nil {
		return err
	}
}
