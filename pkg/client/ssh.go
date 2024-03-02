package client

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHNode runs ssh to a node
func SSHNode(host, port, nodeName, user, keypath string) error {
	node, err := GetNode(host, port, nodeName)
	if err != nil {
		return err
	}
	fmt.Println("Reading key from", keypath)
	key, err := os.ReadFile(keypath)
	if err != nil {
		return err
	}
	fmt.Println("Parsing key")
	sshKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(sshKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
	fmt.Println("Connecting to", node.Host)
	sshClient, err := ssh.Dial("tcp", node.Host+":22", sshConfig)
	if err != nil {
		return err
	}
	defer sshClient.Close()
	fmt.Println("Opening session")
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	fmt.Println("Running bash")
	err = session.Run("bash")
	return err
}

// SSHVM runs ssh to a vm
func SSHVM(host, port, vmName, user string) error {
	return nil
}
