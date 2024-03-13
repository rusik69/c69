package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/rusik69/govnocloud/pkg/types"
)

// UploadFile uploads a file.
func UploadFile(masterHost, masterPort, sourcePath string) error {
	file, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()
	fileName := filepath.Base(sourcePath)
	fileStats, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileStats.Size()
	url := "http://" + masterHost + ":" + masterPort + "/api/v1/files"
	var tempFile types.File
	tempFile.Name = fileName
	tempFile.Size = uint64(fileSize)
	tempFile.Timestamp = uint64(time.Now().Unix())
	tempFileBody, err := json.Marshal(tempFile)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(tempFileBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	var node types.Node
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyText, &node)
	if err != nil {
		return err
	}
	url = "http://" + node.Host + ":" + node.Port + "/api/v1/file/" + fileName
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	url = "http://" + masterHost + ":" + masterPort + "/api/v1/filecommit/" + fileName
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	return nil
}

// DownloadFile downloads a file.
func DownloadFile(masterHost, masterPort, fileName string) error {
	url := "http://" + masterHost + ":" + masterPort + "/api/v1/file/" + fileName
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var file types.File
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyText, &file)
	if err != nil {
		return err
	}
	url = "http://" + file.NodeHost + ":" + file.NodePort + "/api/v1/file/" + fileName
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile deletes a file.
func DeleteFile(host, port, name string) error {
	url := "http://" + host + ":" + port + "/api/v1/file/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyText))
	}
	return nil
}

// ListFiles lists files.
func ListFiles(masterHost, masterPort string) ([]types.File, error) {
	url := "http://" + masterHost + ":" + masterPort + "/api/v1/files"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var files []types.File
	err = json.NewDecoder(resp.Body).Decode(&files)
	if err != nil {
		return nil, err
	}
	return files, nil
}
