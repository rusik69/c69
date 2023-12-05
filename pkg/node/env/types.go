package env

type NodeEnv struct {
	// Name is the name of the node.
	Name string `json:"name"`
	// IP is the IP address of the node.
	IP string `json:"ip"`
	// Port is the port of the node.
	Port            string `json:"port"`
	LibVirtURI      string `json:"libvirt_socket"`
	LibVirtImageDir string `json:"libvirt_image_dir"`
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NodeEnvInstance *NodeEnv
