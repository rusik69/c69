package node

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"encoding/xml"

	"github.com/gin-gonic/gin"

	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

// LibvirtConnection is the singleton instance of libvirt.Connection.
var LibvirtConnection *libvirt.Connect

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
	vm, code, err := CreateVM(tempVM)
	if err != nil {
		logrus.Error("Failed to create VM", err.Error())
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	logrus.Println("Created VM", vm)
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
		c.JSON(500, gin.H{"error": err.Error()})
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

// DownloadImages downloads the images.
func DownloadImages() error {
	for _, img := range types.VMImages {
		if img.Img == "" {
			continue
		}
		imgName := filepath.Join(types.NodeEnvInstance.LibVirtImageDir, img.Img)
		if _, err := os.Stat(imgName); os.IsNotExist(err) {
			err := DownloadFile(img.URL, types.NodeEnvInstance.LibVirtImageDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CreateVM creates the vm.
func CreateVM(vm types.VM) (types.VM, int, error) {
	flavor, ok := types.VMFlavors[vm.Flavor]
	if !ok {
		return types.VM{}, 400, errors.New("flavor not found")
	}
	if types.VMImages[vm.Image].Img == "" {
		return types.VM{}, 400, errors.New(vm.Image + " image not found")
	}
	imgName := filepath.Join(types.NodeEnvInstance.LibVirtImageDir,
		types.VMImages[vm.Image].Img)
	if _, err := os.Stat(imgName); os.IsNotExist(err) {
		return types.VM{}, 400, errors.New(vm.Image + " image not found")
	}
	destImgName := filepath.Join(types.NodeEnvInstance.LibVirtImageDir,
		vm.Name+".qcow2")
	logrus.Println("Copying image", imgName, "to", destImgName)
	err := copyFile(imgName, destImgName)
	if err != nil {
		return types.VM{}, 500, err
	}
	logrus.Println("Resizing image", destImgName, "to", flavor.Disk, "GB")
	err = resizeImage(destImgName, flavor)
	if err != nil {
		return types.VM{}, 500, err
	}
	var cpuShares uint
	var vcpus uint
	if flavor.MilliCPUs > 1024 {
		cpuShares = 1024
		vcpus = uint(flavor.MilliCPUs / 1024)
	} else {
		cpuShares = uint(flavor.MilliCPUs)
		vcpus = 1
	}
	err = PrepareImage(destImgName, vm.Name)
	if err != nil {
		return types.VM{}, 500, err
	}
	domainXML := libvirtxml.Domain{
		Type: "kvm",
		Name: vm.Name,
		Memory: &libvirtxml.DomainMemory{
			Value: uint(flavor.RAM),
			Unit:  "MB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: vcpus,
		},
		CPUTune: &libvirtxml.DomainCPUTune{
			Shares: &libvirtxml.DomainCPUTuneShares{
				Value: cpuShares,
			},
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
						Listen:   types.NodeEnvInstance.ListenHost,
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
							File: destImgName,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda",
						Bus: "virtio",
					},
				},
			},
			Interfaces: []libvirtxml.DomainInterface{
				{
					Source: &libvirtxml.DomainInterfaceSource{
						Network: &libvirtxml.DomainInterfaceSourceNetwork{
							Network: "bridge0",
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
		return types.VM{}, 500, err
	}
	domain, err := LibvirtConnection.DomainDefineXML(vmxml)
	if err != nil {
		logrus.Error(err.Error())
		return types.VM{}, 500, err
	}
	defer func() {
		err := domain.Free()
		if err != nil {
			return
		}
	}()
	err = domain.Create()
	if err != nil {
		return types.VM{}, 500, err
	}
	id, err := domain.GetID()
	if err != nil {
		return types.VM{}, 500, err
	}
	vmDesc, err := domain.GetXMLDesc(libvirt.DomainXMLFlags(0))
	if err != nil {
		return types.VM{}, 500, err
	}
	var vmXML libvirtxml.Domain
	err = xml.Unmarshal([]byte(vmDesc), &vmXML)
	if err != nil {
		return types.VM{}, 500, err
	}
	ip, err := waitForVMUp(domain)
	if err != nil {
		return types.VM{}, 500, err
	}
	err = waitForSSH(ip)
	if err != nil {
		return types.VM{}, 500, err
	}
	err = applyAnsible(ip, "/etc/govnocloud/ansible/vm.yml", vm.Name)
	if err != nil {
		return types.VM{}, 500, err
	}
	var kubeconfig string
	if vm.Image == "k8s" {
		err = applyAnsible(ip, "/etc/govnocloud/ansible/k8s.yml", vm.Name)
		if err != nil {
			return types.VM{}, 500, err
		}
		kubeconfig, err = getKubeConfig(ip)
		if err != nil {
			return types.VM{}, 500, err
		}
		kubeconfig = strings.Replace(kubeconfig, "127.0.0.1", vm.IP, 1)
	}
	vncPort := vmXML.Devices.Graphics[0].VNC.Port
	vncPortString := fmt.Sprintf("%d", vncPort)
	vncURL := "ws://" + types.NodeEnvInstance.IP + ":" + vncPortString
	vm.NodeHostname = types.NodeEnvInstance.IP
	vm.NodePort = types.NodeEnvInstance.ListenPort
	vm.ID = int(id)
	vm.VNCURL = vncURL
	vm.IP = ip
	vm.KubeConfig = kubeconfig
	vm.State = "running"
	return vm, 200, nil
}

// DeleteVM deletes the vm.
func DeleteVM(vm types.VM) error {
	domain, err := LibvirtConnection.LookupDomainByName(vm.Name)
	if err != nil {
		return fmt.Errorf("failed to lookup domain: %w", err)
	}
	defer func() {
		err := domain.Free()
		if err != nil {
			return
		}
	}()
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
	defer func() {
		err := domain.Free()
		if err != nil {
			return
		}
	}()
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
	defer func() {
		err := domain.Free()
		if err != nil {
			return
		}
	}()
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
	for _, domain := range domains {
		defer domain.Free()
	}
	vms := make([]types.VM, 0, len(domains))
	for _, domain := range domains {
		name, err := domain.GetName()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain name: %w", err)
		}
		logrus.Println("VM name", name)
		state, _, err := domain.GetState()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain state: %w", err)
		}
		logrus.Println("VM state", state)
		id, err := domain.GetID()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain id: %w", err)
		}
		logrus.Println("VM id", id)
		vm := types.VM{
			Name:  name,
			State: ParseState(state),
			ID:    int(id),
		}
		vms = append(vms, vm)
	}

	return vms, nil
}
