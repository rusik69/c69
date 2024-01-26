package types

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
	// Host is the Host address of the node.
	Host string `json:"ip"`
	// Port is the port of the node.
	Port string `json:"port"`
}

// File represents a file.
type File struct {
	// Name is the name of the file.
	Name string `json:"name"`
	// Size is the size of the file.
	Size int64 `json:"size"`
	// NodeHost is the node of the file.
	NodeHost string `json:"nodehost"`
	// NodePort is the node of the file.
	NodePort string `json:"nodeport"`
	// NodeName is the node of the file.
	NodeName string `json:"nodename"`
	// Committed is the committed status of the file.
	Committed bool `json:"committed"`
	// Timestamp is the timestamp of the file.
	Timestamp int64 `json:"timestamp"`
}

// MasterEnvInstance is the singleton instance of MasterEnv.
var MasterEnvInstance *MasterEnv
