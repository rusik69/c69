package etcd

import (
	"context"
	"time"
)

// Get gets the value of the key.
func Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := Client.Get(ctx, key)
	cancel()
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}
