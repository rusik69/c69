package node

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

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
func CreateSSHKey() error {
	// check if file exists
	fileName := "/root/.ssh/id_rsa"
	if _, err := os.Stat(fileName); err == nil {
		return nil
	}
	logrus.Println("Creating ssh key", fileName)
	cmd := exec.Command("ssh-keygen", "-f", fileName, "-t", "rsa", "-N", "")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// GetSSHPublicKey gets the ssh public key.
func GetSSHPublicKey() (string, error) {
	file, err := os.Open("/root/.ssh/id_rsa.pub")
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

// AddSSHPublicKey adds the ssh public key to image.
func AddSSHPublicKey(image string, publicKey string) error {
	logrus.Println("Adding ssh public key to", image)
	cmd := exec.Command("qemu-nbd", "-c", "/dev/nbd0", image)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	defer exec.Command("qemu-nbd", "-d", "/dev/nbd0").Run()
	_, err = os.Stat(types.NodeEnvInstance.NbdMountPoint)
	if os.IsNotExist(err) {
		err = os.Mkdir(types.NodeEnvInstance.NbdMountPoint, os.FileMode(0755))
		if err != nil {
			return err
		}
	}
	// wait for /dev/nbd0p1 to appear
	count := 0
	for {
		if count > 100 {
			return errors.New("timeout waiting for /dev/nbd0p1")
		}
		_, err := os.Stat("/dev/nbd0p1")
		if err == nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
		count++
	}
	cmd = exec.Command("mount", "/dev/nbd0p1", types.NodeEnvInstance.NbdMountPoint)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	defer exec.Command("umount", types.NodeEnvInstance.NbdMountPoint).Run()
	// create /nb0/home/ubuntu/.ssh directory recursively
	err = os.MkdirAll(filepath.Join(types.NodeEnvInstance.NbdMountPoint, "/root/.ssh"), os.FileMode(0755))
	if err != nil {
		return err
	}
	authkeysFile := filepath.Join(types.NodeEnvInstance.NbdMountPoint, "/root/.ssh/authorized_keys")
	file, err := os.OpenFile(authkeysFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(publicKey)
	if err != nil {
		return err
	}
	return nil
}
