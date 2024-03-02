package client

import (
	"os"

	"golang.org/x/crypto/ssh"
)

// SSHNode runs ssh to a node
func SSHNode(host, port, nodeName, user, keypath string) error {
	node, err := GetNode(host, port, nodeName)
	if err != nil {
		return err
	}
	key, err := os.ReadFile(keypath)
	if err != nil {
		return err
	}
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
	}
	sshClient, err := ssh.Dial("tcp", node.Host+":22", sshConfig)
	if err != nil {
		return err
	}
	defer sshClient.Close()
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	err = session.Run("bash")
	return err
}

// SSHVM runs ssh to a vm
func SSHVM(host, port, vmName, user string) error {
	return nil
}
