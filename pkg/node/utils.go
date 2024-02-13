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
	return nil
}

// createCloudInit creates the cloud-init iso.
func createCloudInit(vmName, vmType, sshKey, passwordHash string) (string, error) {
	logrus.Println("Creating cloud-init iso")
	filename := types.NodeEnvInstance.LibVirtImageDir + "/" + vmName + "-cloud-init.cfg"
	var userData string
	switch vmType {
	case "ubuntu":
		userData = `#cloud-config
disable_root: false
users:
  - name: work
	shell: /bin/bash
	sudo: ALL=(ALL) NOPASSWD:ALL
	passwd: ` + passwordHash + `
	ssh_authorized_keys:
	  - ssh-rsa ` + sshKey + `
  - name: root
	shell: /bin/bash
	passwd: ` + passwordHash + `
	ssh_authorized_keys:
	  - ssh-rsa ` + sshKey
	}
	userDataFile, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer userDataFile.Close()
	_, err = userDataFile.WriteString(userData)
	if err != nil {
		return "", err
	}
	isoFileName := types.NodeEnvInstance.LibVirtImageDir + "/" + vmName + "-cloud-init.iso"
	cmd := exec.Command("cloud-localds", isoFileName, filename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Println(string(output))
		return "", err
	}
	return isoFileName, nil
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

// CreateSSHKey creates the ssh key.
func CreateSSHKey(fileName string) error {
	// check if file exists
	if _, err := os.Stat(fileName); err == nil {
		return nil
	}
	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-N", "", "-f", fileName)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// GetSSHPublicKey gets the ssh public key.
func GetSSHPublicKey(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
