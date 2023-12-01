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
	Status int `json:"status"`
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

var Flavors = map[string]Flavor{
	"tiny": Flavor{
		ID:    "0",
		Name:  "tiny",
		VCPUs: 1,
		RAM:   512,
		Disk:  2,
	},
	"small": Flavor{
		ID:    "1",
		Name:  "small",
		VCPUs: 1,
		RAM:   1024,
		Disk:  10,
	},
	"medium": Flavor{
		ID:    "2",
		Name:  "medium",
		VCPUs: 2,
		RAM:   2048,
		Disk:  20,
	},
	"large": Flavor{
		ID:    "3",
		Name:  "large",
		VCPUs: 4,
		RAM:   4096,
		Disk:  40,
	},
	"xlarge": Flavor{
		ID:    "4",
		Name:  "xlarge",
		VCPUs: 8,
		RAM:   8192,
		Disk:  80,
	},
}

type Image struct {
	// ID is the ID of the image.
	ID string `json:"id"`
	// Name is the name of the image.
	Img string `json:"img"`
	// URL is the URL of the image.
	URL string `json:"url"`
}

var Images = map[string]Image{
	"ubuntu22.04": Image{
		ID:  "0",
		Img: "ubuntu-22.04-server-cloudimg-amd64.img",
		URL: "https://cloud-images.ubuntu.com/releases/22.04/release/ubuntu-22.04-server-cloudimg-amd64-disk-kvm.img",
	},
	"ubuntu20.04": Image{
		ID:  "1",
		Img: "ubuntu-20.04-server-cloudimg-amd64.img",
		URL: "https://cloud-images.ubuntu.com/releases/20.04/release/ubuntu-20.04-server-cloudimg-amd64-disk-kvm.img",
	},
}
