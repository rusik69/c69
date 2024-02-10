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
	// ListenHost is the IP of the master.
	ListenHost string `json:"ip"`
}

// Node represents a node.
type Node struct {
	// Name is the name of the node.
	Name string `json:"name"`
	// Host is the Host address of the node.
	Host string `json:"ip"`
	// Port is the port of the node.
	Port string `json:"port"`
	// MilliCPUSTotal is the total number of CPUs of the node.
	MilliCPUSTotal uint64 `json:"millicpus_total"`
	// MilliCPUSUsed is the used number of CPUs of the node.
	MilliCPUSUsed uint64 `json:"millicpus_used"`
	// MemoryTotal is the total memory of the node.
	MemoryTotal uint64 `json:"memory_total"`
	// MemoryUsed is the used memory of the node.
	MemoryUsed uint64 `json:"memory_used"`
	// DiskTotal is the total disk of the node.
	DiskTotal uint64 `json:"disk_total"`
	// DiskUsed is the used disk of the node.
	DiskUsed uint64 `json:"disk_used"`
}

// File represents a file.
type File struct {
	// Name is the name of the file.
	Name string `json:"name"`
	// Size is the size of the file.
	Size uint64 `json:"size"`
	// NodeHost is the node of the file.
	NodeHost string `json:"nodehost"`
	// NodePort is the node of the file.
	NodePort string `json:"nodeport"`
	// NodeName is the node of the file.
	NodeName string `json:"nodename"`
	// Committed is the committed status of the file.
	Committed bool `json:"committed"`
	// Timestamp is the timestamp of the file.
	Timestamp uint64 `json:"timestamp"`
}

// MasterEnvInstance is the singleton instance of MasterEnv.
var MasterEnvInstance *MasterEnv
