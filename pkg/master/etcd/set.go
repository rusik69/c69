package etcd

import (
	"context"
	"time"
)

// Put puts the value of the key.
func Put(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := Client.Put(ctx, key, value)
	cancel()
	return err
}
