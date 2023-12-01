package env

import (
	"errors"
	"os"
)

// Parse parses the node environment.
func Parse() (*NodeEnv, error) {
	name := os.Getenv("NODE_NAME")
	if name == "" {
		return nil, errors.New("NODE_NAME is not set")
	}
	id := os.Getenv("NODE_ID")
	if id == "" {
		return nil, errors.New("NODE_ID is not set")
	}
	libvirtURI := os.Getenv("NODE_LIBVIRT_SOCKET")
	if libvirtURI == "" {
		libvirtURI = "qemu:///system"
	}
	ip := os.Getenv("NODE_IP")
	if ip == "" {
		return nil, errors.New("NODE_IP is not set")
	}
	port := os.Getenv("NODE_PORT")
	if port == "" {
		port = "6969"
	}
	libvirtImageDir := os.Getenv("NODE_LIBVIRT_IMAGE_DIR")
	if libvirtImageDir == "" {
		libvirtImageDir = "/var/lib/libvirt/images"
	}
	NodeEnvInstance = &NodeEnv{
		ID:              id,
		Name:            name,
		IP:              ip,
		Port:            port,
		LibVirtURI:      libvirtURI,
		LibVirtImageDir: libvirtImageDir,
	}
	return NodeEnvInstance, nil
}
