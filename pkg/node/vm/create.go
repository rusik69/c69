package vm

import "libvirt.org/go/libvirtxml"

// Create creates the vm.
func (vm VM) Create() error {
	flavor := Flavors[Flavor]

	domain := libvirtxml.Domain{
		Type: "kvm",
		Name: vm.Name,
		Memory: &libvirtxml.DomainMemory{
			Value: flavor.Memory,
			Unit:  "MB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: flavor.VCPU,
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Arch:    "x86_64",
				Machine: "pc",
				Type:    "hvm",
			},
			BootDevices: []libvirtxml.DomainBootDevice{
				{

}
