package types

import "errors"

// Container represents a container.
type Container struct {
	// ID is the ID of the container.
	ID string `json:"id"`
	// Name is the name of the container.
	Name string `json:"name"`
	// Image is the image of the container.
	Image string `json:"image"`
	// State is the state of the container.
	State string `json:"state"`
	// IP is the IP address of the container.
	IP string `json:"ip"`
	// Host is the host of the container.
	Host string `json:"host"`
	// Volumes is the volumes of the container.
	Volumes []Volume `json:"volumes"`
	// Committed is the committed status of the container.
	Committed bool `json:"committed"`
}

// VM represents a virtual machine.
type VM struct {
	// ID is the ID of the virtual machine.
	ID int `json:"id"`
	// Name is the name of the virtual machine.
	Name string `json:"name"`
	// IP is the IP address of the virtual machine.
	IP string `json:"ip"`
	// Node is the host of the virtual machine.
	Node string `json:"host"`
	// NodeHostname is the hostname of the node.
	NodeHostname string `json:"nodehostname"`
	// NodePort is the port of the node.
	NodePort string `json:"nodeport"`
	// State is the status of the virtual machine.
	State string `json:"status"`
	// Image is the image of the virtual machine.
	Image string `json:"image"`
	// Flavor is the flavor of the virtual machine.
	Flavor string `json:"flavor"`
	// Volumes is the volumes of the virtual machine.
	Volumes []Volume `json:"volumes"`
	// Committed is the committed status of the virtual machine.
	Committed bool `json:"committed"`
	// VNCURL is the VNC url of the virtual machine.
	VNCURL string `json:"vncurl"`
}

// Flavor represents a vm flavor.
type VMFlavor struct {
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

var VMFlavors = map[string]VMFlavor{
	"tiny": VMFlavor{
		ID:    "0",
		Name:  "tiny",
		VCPUs: 1,
		RAM:   512,
		Disk:  2,
	},
	"small": VMFlavor{
		ID:    "1",
		Name:  "small",
		VCPUs: 1,
		RAM:   1024,
		Disk:  10,
	},
	"medium": VMFlavor{
		ID:    "2",
		Name:  "medium",
		VCPUs: 2,
		RAM:   2048,
		Disk:  20,
	},
	"large": VMFlavor{
		ID:    "3",
		Name:  "large",
		VCPUs: 4,
		RAM:   4096,
		Disk:  40,
	},
	"xlarge": VMFlavor{
		ID:    "4",
		Name:  "xlarge",
		VCPUs: 8,
		RAM:   8192,
		Disk:  80,
	},
}

type VMImage struct {
	// ID is the ID of the image.
	ID string `json:"id"`
	// Name is the name of the image.
	Img string `json:"img"`
	// URL is the URL of the image.
	URL string `json:"url"`
	// Cmdline is the cmdline of the image.
	Cmdline string `json:"cmdline"`
}

var VMImages = map[string]VMImage{
	/*	"ubuntu22.04": VMImage{
			ID:  "0",
			Img: "ubuntu-22.04.3-live-server-amd64.iso",
			URL: "https://releases.ubuntu.com/jammy/ubuntu-22.04.3-live-server-amd64.iso",
		},
		"ubuntu20.04": VMImage{
			ID:  "1",
			Img: "ubuntu-20.04.6-live-server-amd64.iso",
			URL: "https://releases.ubuntu.com/focal/ubuntu-20.04.6-live-server-amd64.iso",
		},
	*/
	"fedora39": VMImage{
		ID:      "2",
		Img:     "Fedora-Server-netinst-x86_64-39-1.5.iso",
		URL:     "https://download.fedoraproject.org/pub/fedora/linux/releases/39/Server/x86_64/iso/Fedora-Server-netinst-x86_64-39-1.5.iso",
		Cmdline: "ks=file:/var/lib/libvirt/images/fedora39.ks console=ttys0",
	},
}

type NodeEnv struct {
	Name            string `json:"name"`
	IP              string `json:"ip"`
	ListenPort      string `json:"listen_port"`
	ListenHost      string `json:"listen_host"`
	LibVirtURI      string `json:"libvirt_socket"`
	LibVirtImageDir string `json:"libvirt_image_dir"`
	FilesDir        string `json:"files_dir"`
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NodeEnvInstance *NodeEnv

// NodeStats represents the stats.
type NodeStats struct {
	CPUs      int   `json:"cpus"`
	TotalMEM  int64 `json:"total_mem"`
	FreeMEM   int64 `json:"mem"`
	TotalDISK int64 `json:"total_disk"`
	FreeDISK  int64 `json:"disk"`
}

// Volume represents a volume.
type Volume struct {
	// ID is the ID of the volume.
	ID string `json:"id"`
	// Name is the name of the volume.
	Name string `json:"name"`
	// Path is the path of the volume.
	Path string `json:"path"`
	// Size is the size of the volume.
	Size int64 `json:"size"`
}

// findNodeByName finds a node by name.
func FindNodeByName(name string) (Node, error) {
	for _, node := range MasterEnvInstance.Nodes {
		if node.Name == name {
			return node, nil
		}
	}
	return Node{}, errors.New("node " + name + " not found")
}
