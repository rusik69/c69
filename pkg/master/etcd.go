package master

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/rusik69/simplecloud/pkg/types"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ETCDConnect connects to the database.
func ETCDConnect(host, port, user, pass string) (*clientv3.Client, error) {
	var conf clientv3.Config
	if types.MasterEnvInstance.ETCDUser != "" {
		conf = clientv3.Config{
			Endpoints:   []string{"http://" + host + ":" + port},
			DialTimeout: 60 * time.Second,
			Username:    user,
			Password:    pass,
		}
	} else {
		conf = clientv3.Config{
			Endpoints:   []string{"http://" + host + ":" + port},
			DialTimeout: 60 * time.Second,
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

// GetNodes gets the nodes from the database.
func GetNodes() ([]types.Node, error) {
	nodes, err := ETCDList("/nodes/")
	if err != nil {
		return nil, err
	}
	var res []types.Node
	for _, nodeName := range nodes {
		nodeString, err := ETCDGet(nodeName)
		if err != nil {
			return nil, err
		}
		var node types.Node
		err = json.Unmarshal([]byte(nodeString), &node)
		if err != nil {
			return nil, err
		}
		res = append(res, node)
	}
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res, nil
}

// GetNode gets the node from the database.
func GetNode(name string) (types.Node, error) {
	nodeString, err := ETCDGet("/nodes/" + name)
	if err != nil {
		return types.Node{}, err
	}
	if nodeString == "" {
		return types.Node{}, nil
	}
	var node types.Node
	err = json.Unmarshal([]byte(nodeString), &node)
	if err != nil {
		return types.Node{}, err
	}
	return node, nil
}
