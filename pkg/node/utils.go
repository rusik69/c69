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
func createCloudInit(vmName, sshKey string) (string, error) {
	logrus.Println("Creating cloud-init iso")
	filename := types.NodeEnvInstance.LibVirtImageDir + "/" + vmName + "-cloud-init.cfg"
	userData := `#cloud-config
	disable_root: false
	users:
	  - name: work
		shell: /bin/bash
		sudo: true
		passwd: $6$JeZUUZ771KMKfRgI$rConZlL.UqJxCU3VYyimgUun4toLvWQ8LgfxasNwC5XXQgkQsxPgnWrxi8SzI7GO6XhbMHtdrW89s5KIZV1nm0
		ssh_authorized_keys:
		  - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCf743YdT1CjsAIb8ptN2AE2/LOhq2Qp+9o78vpr+l7pfb0Mx140dUxVp+IuqbtTVv2swCyMD8n/sHXND/Yy2T4ekoOQQvuIP0o5UJEFerKr+3HydfctNFb8DOdB/joc7EkdF6a7pqbQGDE9wxPrZphGIhzFCJzQoRjaNUo6JOHd/lJHKc8potHvKJ/ef0mXyHCoEvHaDeragV5SIzozSSeWMUwKR+VgGu/tt/fY6PXr5p596u39CoMkngtGm0ROTLMj/vBHUrcdhMgFrqkzinMxPxR2bw0O9Y9/43s2B/H0abr/YhBxBdFVDlY2msKog5K8cr1vOCF4QCIZUMTHIMOh4uRVVnzPNPvSzCUP5ckotkrnajjG+kc5yNq3qI5PA9UE7twU4unF9T9wBwsYNPkRM1eQbOcs7T5M9DHM6E9PQJZzdTGMLLbiErSfFRbIqz/GFptmrTiFLUrIG7txmRRFW0H04OtfnPwBA6C2v4z7bWaEnRfFlWlxmTaT31APyE= root@x230
	  - name: root
		shell: /bin/bash
		passwd: $6$mYbPgu4O.jOCejHE$.clC6joK06iMMCkNXx0HCtdbNNlmiF1mjwc24l9fGM7Ufcl8/loxD/Nf9F.ap7d6zQUFawtcvhNlKTf2GqxLO/
		ssh_authorized_keys:
		  - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCf743YdT1CjsAIb8ptN2AE2/LOhq2Qp+9o78vpr+l7pfb0Mx140dUxVp+IuqbtTVv2swCyMD8n/sHXND/Yy2T4ekoOQQvuIP0o5UJEFerKr+3HydfctNFb8DOdB/joc7EkdF6a7pqbQGDE9wxPrZphGIhzFCJzQoRjaNUo6JOHd/lJHKc8potHvKJ/ef0mXyHCoEvHaDeragV5SIzozSSeWMUwKR+VgGu/tt/fY6PXr5p596u39CoMkngtGm0ROTLMj/vBHUrcdhMgFrqkzinMxPxR2bw0O9Y9/43s2B/H0abr/YhBxBdFVDlY2msKog5K8cr1vOCF4QCIZUMTHIMOh4uRVVnzPNPvSzCUP5ckotkrnajjG+kc5yNq3qI5PA9UE7twU4unF9T9wBwsYNPkRM1eQbOcs7T5M9DHM6E9PQJZzdTGMLLbiErSfFRbIqz/GFptmrTiFLUrIG7txmRRFW0H04OtfnPwBA6C2v4z7bWaEnRfFlWlxmTaT31APyE=`
	userDataFile, err := os.Create(filename)
	if err != nil {
		return "", nil
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

// createSSHKey creates the ssh key.
func createSSHKey(fileName string) error {
	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-N", "", "-f", fileName)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
