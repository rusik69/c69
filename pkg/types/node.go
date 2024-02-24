package types

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
	// MilliCPUs is the number of MilliCPUs of the flavor.
	MilliCPUs uint64 `json:"millicpus"`
	// RAM is the RAM of the flavor.
	RAM uint64 `json:"ram"`
	// Disk is the disk of the flavor.
	Disk uint64 `json:"disk"`
}

var VMFlavors = map[string]VMFlavor{
	"small": VMFlavor{
		ID:        "1",
		Name:      "small",
		MilliCPUs: 512,
		RAM:       1024,
		Disk:      8,
	},
	"medium": VMFlavor{
		ID:        "2",
		Name:      "medium",
		MilliCPUs: 1024,
		RAM:       2048,
		Disk:      8,
	},
	"large": VMFlavor{
		ID:        "2",
		Name:      "large",
		MilliCPUs: 2048,
		RAM:       4096,
		Disk:      16,
	},
	"xlarge": VMFlavor{
		ID:        "4",
		Name:      "xlarge",
		MilliCPUs: 4096,
		RAM:       8192,
		Disk:      32,
	},
}

type VMImage struct {
	// ID is the ID of the image.
	ID string `json:"id"`
	// Name is the name of the image.
	Img string `json:"img"`
	// URL is the URL of the image.
	URL string `json:"url"`
	// UserDataFile is the cloud init file of the image.
	UserDataFile string `json:"cloudinitfile"`
	// Type is the type of the image.
	Type string `json:"type"`
}

var VMImages = map[string]VMImage{
	"ubuntu22.04": VMImage{
		ID:   "0",
		Type: "ubuntu",
		Img:  "jammy-server-cloudimg-amd64.img",
		URL:  "https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img",
	},
	// "ubuntu20.04": VMImage{
	// 	ID:   "1",
	// 	Type: "ubuntu",
	// 	Img:  "ubuntu-20.04-server-cloudimg-amd64-disk-kvm.img",
	// 	URL:  "https://cloud-images.ubuntu.com/releases/focal/release/ubuntu-20.04-server-cloudimg-amd64-disk-kvm.img",
	// },
	//"fedora39": VMImage{
	//	ID:   "0",
	//		Type: "fedora",
	//	Img:  "Fedora-Server-KVM-39-1.5.x86_64.qcow2",
	//	URL:  "https://download.fedoraproject.org/pub/fedora/linux/releases/39/Server/x86_64/images/Fedora-Server-KVM-39-1.5.x86_64.qcow2",
	//},
}

type NodeEnv struct {
	Name            string `json:"name"`
	IP              string `json:"ip"`
	ListenPort      string `json:"listen_port"`
	ListenHost      string `json:"listen_host"`
	LibVirtURI      string `json:"libvirt_socket"`
	LibVirtImageDir string `json:"libvirt_image_dir"`
	LibVirtBootDir  string `json:"libvirt_boot_dir"`
	FilesDir        string `json:"files_dir"`
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NodeEnvInstance *NodeEnv

// NodeStats represents the stats.
type NodeStats struct {
	MilliCPUs uint64 `json:"millicpus"`
	TotalMEM  uint64 `json:"total_mem"`
	FreeMEM   uint64 `json:"mem"`
	TotalDISK uint64 `json:"total_disk"`
	FreeDISK  uint64 `json:"disk"`
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
