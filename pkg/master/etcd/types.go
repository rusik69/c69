package etcd

import clientv3 "go.etcd.io/etcd/client/v3"

// Client is the database connection.
var Client *clientv3.Client
