package vm

import (
	"github.com/rusik69/govnocloud/pkg/node/env"
	"libvirt.org/go/libvirt"
)

// Connect connects to the libvirt daemon.
func Connect() (*libvirt.Connect, error) {
	conn, err := libvirt.NewConnect(env.NodeEnvInstance.LibVirtURI)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
