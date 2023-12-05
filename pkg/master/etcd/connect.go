package etcd

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/rusik69/govnocloud/pkg/master/env"
)

// Connect connects to the database.
func Connect(host, port, user, pass string) (*clientv3.Client, error) {
	var conf clientv3.Config
	if env.MasterEnvInstance.ETCDUser != "" {
		conf = clientv3.Config{
			Endpoints:   []string{"http://" + host + ":" + port},
			DialTimeout: 10 * time.Second,
			Username:    user,
			Password:    pass,
		}
	} else {
		conf = clientv3.Config{
			Endpoints:   []string{"http://" + host + ":" + port},
			DialTimeout: 10 * time.Second,
		}
	}
	cli, err := clientv3.New(conf)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
