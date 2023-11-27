package vm

import (
	"github.com/rusik69/govnocloud/pkg/node/volume"
	"libvirt.org/go/libvirt"
)

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
	// Status is the status of the virtual machine.
	Status string `json:"status"`
	// Image is the image of the virtual machine.
	Image string `json:"image"`
	// Flavor is the flavor of the virtual machine.
	Flavor string `json:"flavor"`
	// Volumes is the volumes of the virtual machine.
	Volumes []volume.Volume `json:"volumes"`
}

// LibvirtConnection is the singleton instance of libvirt.Connection.
var LibvirtConnection *libvirt.Connect

// Flavor represents a flavor.
type Flavor struct {
	// ID is the ID of the flavor.
	ID string `json:"id"`
	// Name is the name of the flavor.
	Name string `json:"name"`
	// VCPUs is the number of VCPUs of the flavor.
	VCPUs int `json:"vcpus"`
	// RAM is the RAM of the flavor.
	RAM int `json:"ram"`
	// Disk is the disk of the flavor.
	Disk int `json:"disk"`
}

var Flavors = []Flavor{
	Flavor{
		ID:    "0",
		Name:  "tiny",
		VCPUs: 1,
		RAM:   512,
		Disk:  1,
	},
}
