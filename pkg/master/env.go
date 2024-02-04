package master

import (
	"errors"
	"os"
	"strings"

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
		listenPort = "7070"
	}
	nodesString := os.Getenv("MASTER_NODES")
	if nodesString == "" {
		return nil, errors.New("no nodes")
	}
	nodesStringSplit := strings.Split(nodesString, ",")
	var nodes []types.Node
	for _, nodeString := range nodesStringSplit {
		s := strings.Split(nodeString, ":")
		name := s[0]
		host := s[1]
		port := s[2]
		nodes = append(nodes, types.Node{
			Name: name,
			Host: host,
			Port: port,
		})
	}
	return &types.MasterEnv{
		ETCDHost:   etcdhost,
		ETCDPort:   etcdport,
		ETCDUser:   etcduser,
		ETCDPass:   etcdpass,
		ListenPort: listenPort,
		ListenHost: listenHost,
		Nodes:      nodes,
	}, nil
}
