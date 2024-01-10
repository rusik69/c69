package node

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

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
