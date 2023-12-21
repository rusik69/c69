package types

// WEBEnv is the environment of the web server.
type WEBEnv struct {
	// Port is the port of the web server.
	Port string `json:"port"`
	// MasterHost is the host of the master node.
	MasterHost string `json:"master_host"`
	// MasterPort is the port of the master node.
	MasterPort string `json:"master_port"`
}

// WEBEnvInstance is the singleton instance of WEBEnv.
var WEBEnvInstance *WEBEnv
