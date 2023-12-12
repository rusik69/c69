package server

import (
	"context"
	"time"

	"github.com/rusik69/govnocloud/pkg/master/env"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ETCDConnect connects to the database.
func ETCDConnect(host, port, user, pass string) (*clientv3.Client, error) {
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

// ETCDGet gets the value of the key.
func ETCDGet(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := ETCDClient.Get(ctx, key)
	cancel()
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}

// ETCDPut puts the value of the key.
func ETCDPut(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := ETCDClient.Put(ctx, key, value)
	cancel()
	return err
}

// ETCEList lists the keys.
func ETCDList(key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := ETCDClient.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
	}
	return keys, nil
}

// ETCDDelete deletes the key.
func ETCDDelete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := ETCDClient.Delete(ctx, key)
	cancel()
	return err
}

// ETCDClient is the database connection.
var ETCDClient *clientv3.Client
