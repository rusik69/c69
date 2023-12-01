package vm

import "libvirt.org/go/libvirtxml"

// Create creates the vm.
func (vm VM) Create() error {
	flavor := Flavors[vm.Flavor]
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
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: Images[vm.Image].Img,
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
	xml, err := domainXML.Marshal()
	if err != nil {
		return err
	}
	domain, err := LibvirtConnection.DomainDefineXML(xml)
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
