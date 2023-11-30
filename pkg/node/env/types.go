package env

type NodeEnv struct {
	// ID is the ID of the node.
	ID string `json:"id"`
	// Name is the name of the node.
	Name string `json:"name"`
	// IP is the IP address of the node.
	IP string `json:"ip"`
	// Port is the port of the node.
	Port       string `json:"port"`
	LibVirtURI string `json:"libvirt_socket"`
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NodeEnvInstance *NodeEnv
