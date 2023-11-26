package vm

import "github.com/rusik69/govnocloud/pkg/node/volume"

// VM represents a virtual machine.
type VM struct {
	// ID is the ID of the virtual machine.
	ID string `json:"id"`
	// Name is the name of the virtual machine.
	Name string `json:"name"`
	// IP is the IP address of the virtual machine.
	IP string `json:"ip"`
	// Host is the host of the virtual machine.
	Host string `json:"host"`
	// Volumes is the volumes of the virtual machine.
	Volumes []volume.Volume `json:"volumes"`
}
