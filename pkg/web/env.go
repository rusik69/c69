package web

import (
	"os"

	"github.com/rusik69/govnocloud/pkg/types"
)

// Parse parses the environment variables.
func ParseEnv() (*types.WEBEnv, error) {
	port := os.Getenv("WEB_LISTEN_PORT")
	if port == "" {
		port = "8080"
	}
	masterHost := os.Getenv("WEB_MASTER_HOST")
	if masterHost == "" {
		masterHost = "localhost"
	}
	masterPort := os.Getenv("WEB_MASTER_PORT")
	if masterPort == "" {
		masterPort = "7070"
	}
	WEBEnvInstance := &types.WEBEnv{
		Port:       port,
		MasterHost: masterHost,
		MasterPort: masterPort,
	}
	return WEBEnvInstance, nil
}
