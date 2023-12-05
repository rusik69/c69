package vm

import "libvirt.org/go/libvirt"

// Get gets the vm.
func (vm *VM) Get() error {
	domain, err := LibvirtConnection.LookupDomainById(uint32(vm.ID))
	if err != nil {
		return err
	}
	vm.Name, err = domain.GetName()
	if err != nil {
		return err
	}
	state, _, err := domain.GetState()
	if err != nil {
		return err
	}
	switch state {
	case libvirt.DOMAIN_NOSTATE:
		vm.State = "NOSTATE"
	case libvirt.DOMAIN_RUNNING:
		vm.State = "RUNNING"
	case libvirt.DOMAIN_BLOCKED:
		vm.State = "BLOCKED"
	case libvirt.DOMAIN_PAUSED:
		vm.State = "PAUSED"
	case libvirt.DOMAIN_SHUTDOWN:
		vm.State = "SHUTDOWN"
	case libvirt.DOMAIN_SHUTOFF:
		vm.State = "SHUTOFF"
	case libvirt.DOMAIN_CRASHED:
		vm.State = "CRASHED"
	case libvirt.DOMAIN_PMSUSPENDED:
		vm.State = "PMSUSPENDED"
	}
	return nil
}
