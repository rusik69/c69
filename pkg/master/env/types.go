package env

// MasterEnv represents the master environment.
type MasterEnv struct {
	// ETCDHost is the host of the etcd.
	ETCDHost string `json:"etcd_url"`
	// ETCDPort is the port of the etcd.
	ETCDPort string `json:"etcd_port"`
	// ETCDUser is the user of the etcd.
	ETCDUser string `json:"etcd_user"`
	// ETCDPass is the password of the etcd.
	ETCDPass string `json:"etcd_pass"`
	// ListenPort is the port of the master.
	ListenPort string `json:"port"`
	// Nodes is the list of nodes.
	Nodes []Node `json:"nodes"`
}

// Node represents a node.
type Node struct {
	// Name is the name of the node.
	Name string `json:"name"`
	// IP is the IP address of the node.
	IP string `json:"ip"`
	// Port is the port of the node.
	Port string `json:"port"`
}

// MasterEnvInstance is the singleton instance of MasterEnv.
var MasterEnvInstance *MasterEnv
