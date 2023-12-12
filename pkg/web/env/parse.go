package env

import (
	"os"
)

// Parse parses the environment variables.
func Parse() (*WEBEnv, error) {
	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "8080"
	}
	masterHost := os.Getenv("WEB_MASTER_HOST")
	if masterHost == "" {
		masterHost = "localhost"
	}
	masterPort := os.Getenv("WEB_MASTER_PORT")
	if masterPort == "" {
		masterPort = "6969"
	}
	WEBEnvInstance = &WEBEnv{
		Port:       port,
		MasterHost: masterHost,
		MasterPort: masterPort,
	}
	return WEBEnvInstance, nil
}
