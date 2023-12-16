package env

import (
	"os"
)

// Parse parses the environment variables.
func Parse() (*WEBEnv, error) {
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "8080"
	}
	masterHost := os.Getenv("MASTER_HOST")
	if masterHost == "" {
		masterHost = "localhost"
	}
	masterPort := os.Getenv("MASTER_PORT")
	if masterPort == "" {
		masterPort = "7070"
	}
	WEBEnvInstance = &WEBEnv{
		Port:       port,
		MasterHost: masterHost,
		MasterPort: masterPort,
	}
	return WEBEnvInstance, nil
}
