package node

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"encoding/xml"

	"github.com/gin-gonic/gin"

	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

// LibvirtConnection is the singleton instance of libvirt.Connection.
var LibvirtConnection *libvirt.Connect

type Graphics struct {
	Type string `xml:"type,attr"`
	Port int    `xml:"port,attr"`
}

type Devices struct {
	Graphics []Graphics `xml:"graphics"`
}

// VMConnect connects to the libvirt daemon.
func VMConnect() (*libvirt.Connect, error) {
	logrus.Println("Connecting to libvirt daemon at", types.NodeEnvInstance.LibVirtURI)
	conn, err := libvirt.NewConnect(types.NodeEnvInstance.LibVirtURI)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// CreateHandler handles the create request.
func CreateVMHandler(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	var tempVM types.VM
	if err := c.ShouldBindJSON(&tempVM); err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if tempVM.Name == "" || tempVM.Image == "" || tempVM.Flavor == "" {
		logrus.Error("name, image or flavor is empty")
		c.JSON(400, gin.H{"error": "name, image or flavor is empty"})
		return
	}
	logrus.Println("Creating VM", tempVM.Name, tempVM.Image, tempVM.Flavor)
	vm, err := CreateVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logrus.Println("Created VM with id", vm.ID)
	c.JSON(200, vm)
}

// DeleteHandler handles the delete request.
func DeleteVMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	tempVM := types.VM{Name: name}
	logrus.Println("Deleting VM", tempVM)
	err := DeleteVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// ListHandler handles the list request.
func ListVMHandler(c *gin.Context) {
	logrus.Println("Listing VMs")
	vms, err := ListVMs()
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vms)
}

// GetVMHandler handles the get request.
func GetVMHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logrus.Println("Getting VM", idInt)
	tempVM := types.VM{ID: idInt}
	vm, err := GetVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, vm)
}

// StopVMHandler handles the stop vm request.
func StopVMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	tempVM := types.VM{Name: name}
	logrus.Println("Stopping VM", tempVM.ID)
	err := StopVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// StartVMHandler handles the start vm request.
func StartVMHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		logrus.Error("name is empty")
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	tempVM := types.VM{Name: name}
	logrus.Println("Starting VM", tempVM.ID)
	err := StartVM(tempVM)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// CreateVM creates the vm.
func CreateVM(vm types.VM) (types.VM, error) {
	flavor, ok := types.VMFlavors[vm.Flavor]
	if !ok {
		return types.VM{}, errors.New("flavor not found")
	}
	srcFile, err := os.Open(filepath.Join(types.NodeEnvInstance.LibVirtImageDir, types.VMImages[vm.Image].Img))
	if err != nil {
		return types.VM{}, err
	}
	defer srcFile.Close()
	destFileName := filepath.Join(types.NodeEnvInstance.LibVirtImageDir, vm.Name+".qcow2")
	destFile, err := os.Create(destFileName)
	if err != nil {
		return types.VM{}, err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return types.VM{}, err
	}
	err = destFile.Sync()
	if err != nil {
		return types.VM{}, err
	}
	domainXML := libvirtxml.Domain{
		Type: "kvm",
		Name: vm.Name,
		Memory: &libvirtxml.DomainMemory{
			Value: uint(flavor.RAM),
			Unit:  "MB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: uint(flavor.VCPUs),
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Arch:    "x86_64",
				Machine: "pc",
				Type:    "hvm",
			},
			BootDevices: []libvirtxml.DomainBootDevice{
				{
					Dev: "hd",
				},
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Graphics: []libvirtxml.DomainGraphic{
				{
					VNC: &libvirtxml.DomainGraphicVNC{
						AutoPort: "yes",
						Listen:   "0.0.0.0",
					},
				},
			},
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: destFileName,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "vda",
						Bus: "virtio",
					},
				},
			},
			Interfaces: []libvirtxml.DomainInterface{
				{
					Source: &libvirtxml.DomainInterfaceSource{
						Network: &libvirtxml.DomainInterfaceSourceNetwork{
							Network: "default",
						},
					},
					Model: &libvirtxml.DomainInterfaceModel{
						Type: "virtio",
					},
				},
			},
		},
	}
	vmxml, err := domainXML.Marshal()
	if err != nil {
		return types.VM{}, err
	}
	domain, err := LibvirtConnection.DomainDefineXML(vmxml)
	if err != nil {
		logrus.Error(err.Error())
		return types.VM{}, err
	}
	defer domain.Free()
	err = domain.Create()
	if err != nil {
		return types.VM{}, err
	}
	id, err := domain.GetID()
	if err != nil {
		return types.VM{}, err
	}
	vmDesc, err := domain.GetXMLDesc(libvirt.DomainXMLFlags(0))
	if err != nil {
		return types.VM{}, err
	}
	var devices Devices
	err = xml.Unmarshal([]byte(vmDesc), &devices)
	if err != nil {
		fmt.Println("Failed to unmarshal XML")
		return types.VM{}, err
	}
	vncPort := 0
	for _, graphics := range devices.Graphics {
		if graphics.Type == "vnc" {
			vncPort = graphics.Port
		}
	}
	vm.ID = int(id)
	vm.VNCPort = vncPort
	return vm, nil
}

// DeleteVM deletes the vm.
func DeleteVM(vm types.VM) error {
	domain, err := LibvirtConnection.LookupDomainByName(vm.Name)
	if err != nil {
		return fmt.Errorf("failed to lookup domain: %w", err)
	}
	defer domain.Free()

	active, err := domain.IsActive()
	if err != nil {
		return fmt.Errorf("failed to check domain status: %w", err)
	}

	if active {
		err = domain.Destroy()
		if err != nil {
			return fmt.Errorf("failed to destroy domain: %w", err)
		}
	}

	err = domain.Undefine()
	if err != nil {
		return fmt.Errorf("failed to undefine domain: %w", err)
	}
	return nil
}

// StopVM stops the vm.
func StopVM(vm types.VM) error {
	domain, err := LibvirtConnection.LookupDomainByName(vm.Name)
	if err != nil {
		return err
	}
	defer domain.Free()
	err = domain.Destroy()
	if err != nil {
		return err
	}
	return nil
}

// StartVM starts the vm.
func StartVM(vm types.VM) error {
	domain, err := LibvirtConnection.LookupDomainByName(vm.Name)
	if err != nil {
		return err
	}
	defer domain.Free()
	err = domain.Create()
	if err != nil {
		return err
	}
	return nil
}

// GetVM gets the vm.
func GetVM(vm types.VM) (types.VM, error) {
	domain, err := LibvirtConnection.LookupDomainById(uint32(vm.ID))
	if err != nil {
		return types.VM{}, err
	}
	vm.Name, err = domain.GetName()
	if err != nil {
		return types.VM{}, err
	}
	state, _, err := domain.GetState()
	if err != nil {
		return types.VM{}, err
	}
	vm.State = ParseState(state)
	return types.VM{}, nil
}

// ListVMs lists the vms.
func ListVMs() ([]types.VM, error) {
	domains, err := LibvirtConnection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		return nil, fmt.Errorf("failed to list domains: %w", err)
	}
	defer func() {
		for _, domain := range domains {
			domain.Free()
		}
	}()

	vms := make([]types.VM, 0, len(domains))
	for _, domain := range domains {
		name, err := domain.GetName()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain name: %w", err)
		}
		state, _, err := domain.GetState()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain state: %w", err)
		}
		id, err := domain.GetID()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain id: %w", err)
		}
		vm := types.VM{
			Name:  name,
			State: ParseState(state),
			ID:    int(id),
		}
		vms = append(vms, vm)
	}

	return vms, nil
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

// DownloadVMImages downloads the images.
func DownloadVMImages() error {
	for _, image := range types.VMImages {
		_, err := os.Stat(filepath.Join(types.NodeEnvInstance.LibVirtImageDir, path.Base(image.URL)))
		if err != nil && os.IsNotExist(err) {
			logrus.Println("Downloading image", image.URL)
			err := DownloadFile(image.URL, types.NodeEnvInstance.LibVirtImageDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
