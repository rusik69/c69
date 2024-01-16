package client_test

import (
	"os"
	"testing"

	"github.com/rusik69/govnocloud/pkg/client"
)

func TestFileUpload(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tempFile.Name())
	err = client.UploadFile(masterHost, masterPort, tempFile.Name())
	if err != nil {
		t.Error(err)
	}
}
