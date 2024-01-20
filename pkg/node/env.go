package node

import (
	"errors"
	"os"

	"github.com/rusik69/govnocloud/pkg/types"
)

// ParseEnv parses the node environment.
func ParseEnv() (*types.NodeEnv, error) {
	name := os.Getenv("NODE_NAME")
	if name == "" {
		return nil, errors.New("NODE_NAME is not set")
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
	podmanSocket := os.Getenv("NODE_PODMAN_SOCKET")
	if podmanSocket == "" {
		podmanSocket = "unix:/run/podman/podman.sock"
	}
	libvirtImageDir := os.Getenv("NODE_LIBVIRT_IMAGE_DIR")
	if libvirtImageDir == "" {
		libvirtImageDir = "/var/lib/libvirt/images"
	}
	filesDir := os.Getenv("NODE_FILES_DIR")
	if filesDir == "" {
		filesDir = "/mnt"
	}
	NodeEnvInstance := &types.NodeEnv{
		Name:            name,
		IP:              ip,
		Port:            port,
		LibVirtURI:      libvirtURI,
		LibVirtImageDir: libvirtImageDir,
		FilesDir:        filesDir,
		PodmanSocket:    podmanSocket,
	}
	return NodeEnvInstance, nil
}
