package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rusik69/simplecloud/pkg/deploy"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var nodes []string
var master, ansibleInventoryFile string
var key, user, nodesString string

var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a simple cloud",
	Long:  `deploy a simple cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		nodes = strings.Split(nodesString, ",")
		if len(nodes) == 0 || master == "" {
			logrus.Println("Nodes and master must be specified")
			os.Exit(1)
		}
		if nodes[0] == "" {
			logrus.Println("Nodes must be specified")
			os.Exit(1)
		}
		nodesString := strings.Join(nodes, ",")
		logrus.Println("Deploying simplecloud on nodes", nodesString, "and master", master)
		logrus.Println("Generating Ansible inventory file", ansibleInventoryFile)
		err := deploy.GenerateAnsibleConfig(nodes, master, ansibleInventoryFile)
		if err != nil {
			panic(err)
		}
		logrus.Println("Running Ansible on inventory file", ansibleInventoryFile)
		err = deploy.RunAnsible(ansibleInventoryFile)
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			logrus.Println("Stopping simplecloud on node", node)
			err := deploy.RunSSHCommand(node, key, user, "sudo systemctl stop simplecloud-node; cleanup.sh")
			if err != nil {
				panic(err)
			}
		}
		logrus.Println("Stopping simplecloud on master", master)
		err = deploy.RunSSHCommand(master, key, user, "sudo systemctl stop simplecloud-master; cleanup.sh")
		if err != nil {
			panic(err)
		}
		logrus.Println("Running cleanup.sh on master", master)
		err = deploy.RunSSHCommand(master, key, user, "cleanup.sh")
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			logrus.Println("Running cleanup.sh on node", node)
			err := deploy.RunSSHCommand(node, key, user, "cleanup.sh")
			if err != nil {
				panic(err)
			}
		}
		logrus.Println("Copying simplecloud-master-linux-amd64 to master", master)
		err = deploy.CopyFile(master, key, user, "bin/simplecloud-master-linux-amd64", "/usr/local/bin/simplecloud-master")
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			logrus.Println("Copying simplecloud-node-linux-amd64 to node", node)
			err := deploy.CopyFile(node, key, user, "bin/simplecloud-node-linux-amd64", "/usr/local/bin/simplecloud-node")
			if err != nil {
				panic(err)
			}
			err = deploy.SyncDir(node, user, "deployments/ansible", "/var/lib/libvirt/")
			if err != nil {
				panic(err)
			}
		}
		logrus.Println("Starting simplecloud on master", master)
		err = deploy.RunSSHCommand(master, key, user, "sudo systemctl start simplecloud-master")
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			logrus.Println("Starting simplecloud on node", node)
			err := deploy.RunSSHCommand(node, key, user, "sudo systemctl start simplecloud-node")
			if err != nil {
				panic(err)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	currentUserHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	rootCmd.PersistentFlags().StringVar(&nodesString, "nodes", "", "nodes to deploy")
	rootCmd.PersistentFlags().StringVar(&master, "master", "", "master to deploy")
	rootCmd.PersistentFlags().StringVar(&ansibleInventoryFile, "inv", "./deployments/ansible/inventories/deploy_hosts", "ansible inventory file")
	rootCmd.PersistentFlags().StringVar(&key, "key", filepath.Join(currentUserHomeDir, ".ssh/id_rsa"), "private ssh key path")
	rootCmd.PersistentFlags().StringVar(&user, "user", "root", "ssh user")
}

func main() {
	Execute()
}
