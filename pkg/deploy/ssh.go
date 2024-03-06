package deploy

import (
	"os"
	"os/exec"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// RunSSHCommand stops a service on remote vm via ssh
func RunSSHCommand(host, keyPath, user, cmd string) error {
	// Read the private key file
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return err
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	// Configure the SSH client
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // This is insecure; use ssh.FixedHostKey in production
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return err
	}
	defer client.Close()

	// Run the command
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	err = session.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

// CopyFile copies file via ssh
func CopyFile(host, keyPath, user, src, dst string) error {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return err
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	// Configure the SSH client
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // This is insecure; use ssh.FixedHostKey in production
	}
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return err
	}
	defer client.Close()
	sftp, err := sftp.NewClient(client)
	if err != nil {
		return err
	}
	defer sftp.Close()
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := sftp.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	if _, err := dstFile.ReadFrom(srcFile); err != nil {
		return err
	}
	return nil
}

// SyncDir syncs directory via ssh
func SyncDir(host, user, src, dst string) error {
	cmd := "rsync -a " + src + " " + user + "@" + host + ":" + dst
	_, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
