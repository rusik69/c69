package master

import (
	"os"

	"github.com/rusik69/govnocloud/pkg/types"
)

// ParseEnv parses the master environment.
func ParseEnv() (*types.MasterEnv, error) {
	etcdhost := os.Getenv("MASTER_ETCD_HOST")
	if etcdhost == "" {
		etcdhost = "localhost"
	}
	etcdport := os.Getenv("MASTER_ETCD_PORT")
	if etcdport == "" {
		etcdport = "2379"
	}
	etcduser := os.Getenv("MASTER_ETCD_USER")
	if etcduser == "" {
		etcduser = ""
	}
	etcdpass := os.Getenv("MASTER_ETCD_PASS")
	if etcdpass == "" {
		etcdpass = ""
	}
	listenHost := os.Getenv("MASTER_LISTEN_HOST")
	if listenHost == "" {
		listenHost = "127.0.0.1"
	}
	listenPort := os.Getenv("MASTER_LISTEN_PORT")
	if listenPort == "" {
		listenPort = "6969"
	}
	return &types.MasterEnv{
		ETCDHost:   etcdhost,
		ETCDPort:   etcdport,
		ETCDUser:   etcduser,
		ETCDPass:   etcdpass,
		ListenPort: listenPort,
		ListenHost: listenHost,
	}, nil
}
