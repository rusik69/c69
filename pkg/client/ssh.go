package client

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// RunSSH runs ssh to a node or a vm
func RunSSH(host, keyPath, user, proxyHost string) error {
	key, err := os.ReadFile(keyPath)
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
	var cli *ssh.Client
	if proxyHost == "" {
		cli, err = ssh.Dial("tcp", host+":22", sshConfig)
		if err != nil {
			return err
		}
		defer cli.Close()
	} else {
		proxy, err := ssh.Dial("tcp", proxyHost+":22", sshConfig)
		if err != nil {
			return err
		}
		defer proxy.Close()
		conn, err := proxy.Dial("tcp", host+":22")
		if err != nil {
			return err
		}
		defer conn.Close()
		ncc, chans, reqs, err := ssh.NewClientConn(conn, host, sshConfig)
		if err != nil {
			return err
		}
		cli = ssh.NewClient(ncc, chans, reqs)
		defer cli.Close()
	}
	fmt.Println("Opening session")
	session, err := cli.NewSession()
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
}
