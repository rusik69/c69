package vm

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/rusik69/govnocloud/pkg/node/env"
	"github.com/sirupsen/logrus"
	"libvirt.org/go/libvirt"
)

// ParseState parses the state of the vm.
func ParseState(state libvirt.DomainState) string {
	switch state {
	case libvirt.DOMAIN_NOSTATE:
		return "NOSTATE"
	case libvirt.DOMAIN_RUNNING:
		return "RUNNING"
	case libvirt.DOMAIN_BLOCKED:
		return "BLOCKED"
	case libvirt.DOMAIN_PAUSED:
		return "PAUSED"
	case libvirt.DOMAIN_SHUTDOWN:
		return "SHUTDOWN"
	case libvirt.DOMAIN_SHUTOFF:
		return "SHUTOFF"
	case libvirt.DOMAIN_CRASHED:
		return "CRASHED"
	case libvirt.DOMAIN_PMSUSPENDED:
		return "PMSUSPENDED"
	}
	return ""
}

// Download downloads the file.
func Download(url string, dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fileName := path.Base(url)
	filePath := path.Join(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	bufferSize := 4096
	buffer := make([]byte, bufferSize)
	_, err = io.CopyBuffer(file, resp.Body, buffer)
	if err != nil {
		return err
	}
	return err
}

// DownloadImages downloads the images.
func DownloadImages() error {
	for _, image := range Images {
		_, err := os.Stat(filepath.Join(env.NodeEnvInstance.LibVirtImageDir, path.Base(image.URL)))
		if err != nil && os.IsNotExist(err) {
			logrus.Println("Downloading image", image.URL)
			err := Download(image.URL, env.NodeEnvInstance.LibVirtImageDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
