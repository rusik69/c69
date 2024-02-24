package client

import (
	"os"

	"golang.org/x/crypto/ssh"
)

// SSHNode runs ssh to a node
func SSHNode(host, user string) error {
	key, err := os.ReadFile("~/.ssh/id_rsa")
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
	sshClient, err := ssh.Dial("tcp", host+":22", sshConfig)
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
func SSHVM(host, user string) error {
	return nil
}
