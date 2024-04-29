package deploy

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// GenerateAnsibleConfig generates ansible config
func GenerateAnsibleConfig(nodes []string, master, invFile string) error {
	file, err := os.Create(invFile)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString("[all]\n")
	if err != nil {
		return err
	}
	for _, node := range nodes {
		_, err = file.WriteString(node + " ansible_become=true\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString(master + " ansible_become=true\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("[nodes]\n")
	if err != nil {
		return err
	}
	for _, node := range nodes {
		_, err = file.WriteString(node + " ansible_become=true\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("[master]\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString(master + " ansible_become=true\n")
	if err != nil {
		return err
	}
	file.Sync()
	fileContent, err := os.ReadFile(invFile)
	if err != nil {
		return err
	}
	fmt.Println(string(fileContent))
	return nil
}

// RunAnsible runs ansible
func RunAnsible(invFile string) error {
	tailscale_authkey := os.Getenv("TAILSCALE_AUTH_KEY")
	tailscale_accesstoken := os.Getenv("TAILSCALE_ACCESS_TOKEN")
	cmd := exec.Command("ansible-playbook", "-i", invFile, "deployments/ansible/main.yml")
	cmd.Env = append(cmd.Env, "ANSIBLE_HOST_KEY_CHECKING=False")
	cmd.Env = append(cmd.Env, "TAILSCALE_AUTH_KEY="+tailscale_authkey)
	cmd.Env = append(cmd.Env, "TAILSCALE_ACCESS_TOKEN="+tailscale_accesstoken)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = cmd.Wait()
	return err
}
