package env

import (
	"errors"
	"os"
)

// Parse parses the master environment.
func Parse() (*MasterEnv, error) {
	etcdhost := os.Getenv("ETCD_HOST")
	if etcdhost == "" {
		etcdhost = "localhost"
	}
	etcdport := os.Getenv("ETCD_PORT")
	if etcdport == "" {
		etcdport = "2379"
	}
	etcduser := os.Getenv("ETCD_USER")
	if etcduser == "" {
		etcduser = "root"
	}
	etcdpass := os.Getenv("ETCD_PASS")
	if etcdpass == "" {
		etcdpass = "password"
	}
	listenport := os.Getenv("LISTEN_PORT")
	if listenport == "" {
		listenport = "6969"
	}
	nodesString := os.Getenv("NODES")
	if nodesString == "" {
		return nil, errors.New("no nodes")
	}
	nodesStringSplit := strings.Split(nodesString, ",")
	var nodes []Node
	for _, nodeString := range nodesStringSplit {
		s := strings.Split(nodeString, ":")
		name := s[0]
		host := s[1]
		port := s[2]
		nodes = append(nodes, Node{
			Name: name,
			IP:   host,
			Port: port,
		})
	return &MasterEnv{
		ETCDHost:   etcdhost,
		ETCDPort:   etcdport,
		ETCDUser:   etcduser,
		ETCDPass:   etcdpass,
		ListenPort: listenport,
		Nodes:      nodes,
	}, nil
}
