package client_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rusik69/simplecloud/pkg/client"
)

var tempFileName string

func TestFileUpload(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tempFile.Name())
	_, err = tempFile.WriteString("test")
	if err != nil {
		t.Error(err)
	}
	tempFileName = tempFile.Name()
	err = client.UploadFile(masterHost, masterPort, tempFileName)
	if err != nil {
		t.Error(err)
	}
}

func TestFileDownload(t *testing.T) {
	tempFileNameOnly := filepath.Base(tempFileName)
	err := client.DownloadFile(masterHost, masterPort, tempFileNameOnly)
	if err != nil {
		t.Error(err)
	}
}

func TestFileList(t *testing.T) {
	files, err := client.ListFiles(masterHost, masterPort)
	if err != nil {
		t.Error(err)
	}
	if len(files) != 1 {
		t.Error("expected 1 file, got ", len(files))
	}
}

func TestFileDelete(t *testing.T) {
	tempFileNameOnly := filepath.Base(tempFileName)
	err := client.DeleteFile(masterHost, masterPort, tempFileNameOnly)
	if err != nil {
		t.Error(err)
	}
}

// UploadFiles uploads files.
func UploadFiles() {
	for i := 0; i < 10; i++ {
		tempFile, _ := os.CreateTemp("", "testfile-*.txt")
		defer os.Remove(tempFile.Name())
		_, err := tempFile.WriteString("test")
		if err != nil {
			panic(err)
		}
		tempFileName = tempFile.Name()
		err = client.UploadFile(masterHost, masterPort, tempFileName)
		if err != nil {
			panic(err)
		}
	}
}
