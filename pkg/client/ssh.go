package client

import (
	"bufio"
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
	in, _ := session.StdinPipe()
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return err
	}
	if err := session.Shell(); err != nil {
		return err
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)
	}
	return err
}

// SSHVM runs ssh to a vm
func SSHVM(host, port, vmName, user string) error {
	return nil
}
