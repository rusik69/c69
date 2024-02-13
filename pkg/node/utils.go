package node

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
	"libvirt.org/go/libvirt"
)

// resizeImage resizes the image.
func resizeImage(image string, flavor types.VMFlavor) error {
	logrus.Println("Resizing image", image, "to", flavor.Disk, "GB")
	imgInfo, err := exec.Command("qemu-img", "info", image).CombinedOutput()
	if err != nil {
		return err
	}
	lines := strings.Split(string(imgInfo), "\n")
	re := regexp.MustCompile(`^virtual size: (\d+)`)
	var virtualSize string
	for _, line := range lines {
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)
			virtualSize = matches[1]
			break
		}
	}
	size, err := strconv.Atoi(virtualSize)
	if err != nil {
		return err
	}
	var cmdStrings []string
	if size < int(flavor.Disk) {
		cmdStrings = []string{"resize", image, strconv.Itoa(int(flavor.Disk)) + "G"}
	} else {
		cmdStrings = []string{"resize", "--shrink", image, strconv.Itoa(int(flavor.Disk)) + "G"}
	}
	cmd := exec.Command("qemu-img", cmdStrings...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Println(string(output))
		return err
	}
	logrus.Println(string(output))
	return nil
}

// createCloudInit creates the cloud-init iso.
func createCloudInit(vmName, sshKey string) (string, error) {
	logrus.Println("Creating cloud-init iso")
	filename := types.NodeEnvInstance.LibVirtImageDir + "/" + vmName + "-cloud-init.iso"
	userData := `#cloud-config
	hostname: ` + vmName + `
	manage_etc_hosts: true
	users:
	- name: ubuntu
	  passwd: $6$rounds=4096$4b24w.B0WMskMhB/$De/LwoWFnLGTOMYYLpM0lNe8UFPKqk9eU.sZsncaM1StkpAj5w6zgYDRFk6XFW2x8FMdnYwoLwCHlkBTtlEHK1
	  sudo: ALL=(ALL) NOPASSWD:ALL
	  groups: sudo, admin
	  home: /home/ubuntu
	  shell: /bin/bash
	  ssh-authorized-keys:
	  - ` + sshKey
	userDataFile, err := os.Create(filename)
	if err != nil {
		return filename, nil
	}
	defer userDataFile.Close()
	_, err = userDataFile.WriteString(userData)
	if err != nil {
		return filename, err
	}
	return filename, nil
}

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

// DownloadFile downloads the file.
func DownloadFile(url string, dir string) error {
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

// createSSHKey creates the ssh key.
func createSSHKey(fileName string) error {
	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-N", "", "-f", fileName)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
