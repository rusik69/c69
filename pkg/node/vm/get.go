package vm

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
	vm.State = ParseState(state)
	return nil
}
