package vm

import "libvirt.org/go/libvirt"

// Connect connects to the libvirt daemon.
func Connect() (*libvirt.Connect, error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
