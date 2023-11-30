package vm

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

// List lists the vms.
func List() ([]VM, error) {
	domains, err := LibvirtConnection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		return nil, fmt.Errorf("failed to list domains: %w", err)
	}
	defer func() {
		for _, domain := range domains {
			domain.Free()
		}
	}()

	vms := make([]VM, 0, len(domains))
	for _, domain := range domains {
		name, err := domain.GetName()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain name: %w", err)
		}

		state, _, err := domain.GetState()
		if err != nil {
			return nil, fmt.Errorf("failed to get domain state: %w", err)
		}

		vm := VM{
			Name:   name,
			Status: int(state),
		}
		vms = append(vms, vm)
	}

	return vms, nil
}
