package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govnocloud-client",
	Short: "govnocloud is a shitty cloud",
	Long:  `govnocloud is a shitty cloud`,
}

var clientHost string
var clientPort string
var name, image, flavor string
var nodehost, nodeport string
var user, key string
var id, src string
var model, input string

// sshClientCmd represents the ssh commands
var sshClientCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh to vm or node",
	Long:  `ssh to vm or node`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: ssh [vm|node]")
	},
}

// k8sClientCmd represents the ssh commands
var k8sClientCmd = &cobra.Command{
	Use:   "k8s",
	Short: "manage k8s clusters",
	Long:  `manage k8s clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: k8s [create|get|list|delete|stop|start|get-kubeconfig]")
	},
}

// llmClientCmd represents the llm commands
var llmClientCmd = &cobra.Command{
	Use:   "llm",
	Short: "manage llm clusters",
	Long:  `manage llm clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: llm [create|get|list|delete|stop|start]")
	},
}

// k8sCreateCmd represents the k8s create command
var k8sCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create k8s cluster",
	Long:  `create k8s cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		if flavor == "" {
			panic("flavor is required")
		}
		id, err := client.CreateK8S(clientHost, clientPort, name, flavor)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println("K8S cluster created with id " + fmt.Sprint(id))
	},
}

// k8sDeleteCmd represents the k8s delete command
var k8sDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete k8s cluster",
	Long:  `delete k8s cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.DeleteK8S(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// k8sGetCmd represents the k8s get command
var k8sGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get k8s cluster",
	Long:  `get k8s cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		k8s, err := client.GetK8S(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(
			"ID: " + fmt.Sprint(k8s.ID) + "\n" +
				"Name: " + k8s.Name + "\n" +
				"VM: " + k8s.VM.Name + "\n" +
				"Flavor: " + k8s.Flavor + "\n",
		)
	},
}

// k8sListCmd represents the k8s list command
var k8sListCmd = &cobra.Command{
	Use:   "list",
	Short: "list k8s clusters",
	Long:  `list k8s clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		k8sList, err := client.ListK8S(clientHost, clientPort)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(k8sList)
	},
}

// k8sStopCmd represents the k8s stop command
var k8sStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop k8s cluster",
	Long:  `stop k8s cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.StopK8S(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// k8sStartCmd represents the k8s start command
var k8sStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start k8s cluster",
	Long:  `start k8s cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.StartK8S(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// k8sGetKubeconfigCmd represents the k8s get kubeconfig command
var k8sGetKubeconfigCmd = &cobra.Command{
	Use:   "get-kubeconfig",
	Short: "get k8s kubeconfig",
	Long:  `get k8s kubeconfig`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		kubeconfig, err := client.GetKubeconfig(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(kubeconfig)
	},
}

// vmClientCmd represents the vm commands
var vmClientCmd = &cobra.Command{
	Use:   "vm",
	Short: "vm commands",
	Long:  `vm commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: vm [create|delete|list|get]")
	},
}

// fileClientCmd represents the file commands
var fileClientCmd = &cobra.Command{
	Use:   "file",
	Short: "file commands",
	Long:  `file commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: file [upload|download|delete|list]")
	},
}

// nodeCmd represents the node commands
var nodeClientCmd = &cobra.Command{
	Use:   "node",
	Short: "node commands",
	Long:  `node commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: node [add|delete|list|get]")
	},
}

// containerClientCmd represents the container commands
var containerClientCmd = &cobra.Command{
	Use:   "container",
	Short: "container commands",
	Long:  `container commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: container [create|delete|list|get]")
	},
}

// sshNodeCmd represents the ssh node command
var sshNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "ssh to node",
	Long:  `ssh to node`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("node is required")
		}
		nodeName := args[0]
		if key == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				logrus.Panic(err)
			}
			keyPath := filepath.Join(homeDir, ".ssh/id_rsa")
			key = keyPath
		}
		node, err := client.GetNode(clientHost, clientPort, nodeName)
		if err != nil {
			logrus.Panic(err)
		}
		err = client.RunSSH(node.Host, key, user, "")
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// sshVMCmd represents the ssh vm command
var sshVMCmd = &cobra.Command{
	Use:   "vm",
	Short: "ssh to vm",
	Long:  `ssh to vm`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("vm is required")
		}
		vmName := args[0]
		if key == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				logrus.Panic(err)
			}
			keyPath := filepath.Join(homeDir, ".ssh/id_rsa")
			key = keyPath
		}
		vm, err := client.GetVM(clientHost, clientPort, vmName)
		if err != nil {
			logrus.Panic(err)
		}
		err = client.RunSSH(vm.IP, key, user, vm.NodeHostname)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// nodeAddCmd represents the node add command
var nodeAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add node",
	Long:  `add node`,
	Run: func(cmd *cobra.Command, args []string) {
		if nodehost == "" {
			panic("nodehost is required")
		}
		if nodeport == "" {
			panic("nodeport is required")
		}
		if name == "" {
			panic("name is required")
		}
		err := client.AddNode(clientHost, clientPort, name, nodehost, nodeport)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// nodeDeleteCmd represents the node delete command
var nodeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete node",
	Long:  `delete node`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.DeleteNode(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// nodeListCmd represents the node list command
var nodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "list nodes",
	Long:  `list nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		nodes, err := client.ListNodes(clientHost, clientPort)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(nodes)
	},
}

// nodeGetCmd represents the node get command
var nodeGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get node",
	Long:  `get node`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		node, err := client.GetNode(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(
			"Name: " + node.Name + "\n" +
				"Host: " + node.Host + "\n" +
				"Port: " + node.Port + "\n",
		)
	},
}

// vmCreateCmd represents the vm create command
var vmCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create vm",
	Long:  `create vm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		if image == "" {
			panic("image is required")
		}
		if flavor == "" {
			panic("flavor is required")
		}
		id, err := client.CreateVM(clientHost, clientPort, name, image, flavor)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println("VM created with id " + fmt.Sprint(id))
	},
}

// llmCreateCmd represents the llm create command
var llmCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create llm",
	Long:  `create llm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		if model == "" {
			panic("model is required")
		}
		id, err := client.CreateLLM(clientHost, clientPort, name, model)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println("LLM created with id " + fmt.Sprint(id))
	},
}

// llmDeleteCmd represents the llm delete command
var llmDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete llm",
	Long:  `delete llm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.DeleteLLM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// llmGetCmd represents the llm get command
var llmGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get llm",
	Long:  `get llm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		llm, err := client.GetLLM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(
			"Name: " + llm.Name + "\n" +
				"Model: " + llm.Model + "\n",
		)
	},
}

// llmListCmd represents the llm list command
var llmListCmd = &cobra.Command{
	Use:   "list",
	Short: "list llms",
	Long:  `list llms`,
	Run: func(cmd *cobra.Command, args []string) {
		llms, err := client.ListLLMs(clientHost, clientPort)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(llms)
	},
}

// llmStopCmd represents the llm stop command
var llmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop llm",
	Long:  `stop llm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.StopLLM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// llmStartCmd represents the llm start command
var llmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start llm",
	Long:  `start llm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.StartLLM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// llmGenerateCmd represents the llm generate command
var llmGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate llm response",
	Long:  `generate llm response`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		if input == "" {
			panic("input is required")
		}
		res, err := client.GenerateLLM(clientHost, clientPort, name, input)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(res)
	},
}

// vmDeleteCmd represents the vm delete command
var vmDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete vm",
	Long:  `delete vm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.DeleteVM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

var vmGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get vm",
	Long:  `get vm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		vm, err := client.GetVM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(
			"ID: " + fmt.Sprint(vm.ID) + "\n" +
				"Name: " + vm.Name + "\n" +
				"IP: " + vm.IP + "\n" +
				"Host: " + vm.Node + "\n" +
				"State: " + vm.State + "\n" +
				"Image: " + vm.Image + "\n" +
				"Flavor: " + vm.Flavor + "\n",
		)
	},
}

// vmListCmd represents the vm list command
var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "list vm",
	Long:  `list vm`,
	Run: func(cmd *cobra.Command, args []string) {
		vms, err := client.ListVMs(clientHost, clientPort)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(vms)
	},
}

// vmStopCmd represents the vm stop command
var vmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop vm",
	Long:  `stop vm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.StopVM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// vmStartCmd represents the vm start command
var vmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start vm",
	Long:  `start vm`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.StartVM(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// containerCreateCmd represents the container create command
var containerCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create container",
	Long:  `create container`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		if image == "" {
			panic("image is required")
		}
		if flavor == "" {
			panic("flavor is required")
		}
		id, err := client.CreateContainer(clientHost, clientPort, name, image, flavor)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println("Container created with id " + fmt.Sprint(id))
	},
}

// containerDeleteCmd represents the container delete command
var containerDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete container",
	Long:  `delete container`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			panic("id is required")
		}
		err := client.DeleteContainer(clientHost, clientPort, id)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// containerGetCmd represents the container get command
var containerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get container",
	Long:  `get container`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			panic("id is required")
		}
		container, err := client.GetContainer(clientHost, clientPort, id)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(
			"ID: " + fmt.Sprint(container.ID) + "\n" +
				"Name: " + container.Name + "\n" +
				"IP: " + container.IP + "\n" +
				"Host: " + container.Node + "\n" +
				"State: " + container.State + "\n" +
				"Image: " + container.Image + "\n",
		)
	},
}

// containerListCmd represents the container list command
var containerListCmd = &cobra.Command{
	Use:   "list",
	Short: "list container",
	Long:  `list container`,
	Run: func(cmd *cobra.Command, args []string) {
		containers, err := client.ListContainers(clientHost, clientPort)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Printf("| %-10s | %-10s | %-16s | %-5s | %-7s | %-20s |\n", "ID", "NAME", "IP", "Host", "Status", "Image")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		for _, container := range containers {
			fmt.Printf("| %-10s | %-10s | %-16s | %-5s | %-7s | %-20s |\n", container.ID, container.Name, container.IP, container.Node, container.State, container.Image)
		}
	},
}

// containerStopCmd represents the container stop command
var containerStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop container",
	Long:  `stop container`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			panic("id is required")
		}
		err := client.StopContainer(clientHost, clientPort, id)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// containerStartCmd represents the container start command
var containerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start container",
	Long:  `start container`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			panic("id is required")
		}
		err := client.StartContainer(clientHost, clientPort, id)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// fileUploadCmd represents the file upload command
var fileUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file",
	Long:  `upload file`,
	Run: func(cmd *cobra.Command, args []string) {
		srcPtr := cmd.PersistentFlags().String("src", "", "file source")
		if *srcPtr == "" {
			panic("src is required")
		}
		err := client.UploadFile(clientHost, clientPort, *srcPtr)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// fileDownloadCmd represents the file download command
var fileDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download file",
	Long:  `download file`,
	Run: func(cmd *cobra.Command, args []string) {
		if src == "" {
			panic("src is required")
		}
		err := client.DownloadFile(clientHost, clientPort, src)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// fileDeleteCmd represents the file delete command
var fileDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete file",
	Long:  `delete file`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("name is required")
		}
		err := client.DeleteFile(clientHost, clientPort, name)
		if err != nil {
			logrus.Panic(err)
		}
	},
}

// fileListCmd represents the file list command
var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "list files",
	Long:  `list files`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := client.ListFiles(clientHost, clientPort)
		if err != nil {
			logrus.Panic(err)
		}
		fmt.Println(files)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(sshClientCmd)
	rootCmd.AddCommand(vmClientCmd)
	rootCmd.AddCommand(nodeClientCmd)
	rootCmd.AddCommand(containerClientCmd)
	rootCmd.AddCommand(fileClientCmd)
	rootCmd.AddCommand(k8sClientCmd)
	rootCmd.AddCommand(llmClientCmd)
	llmClientCmd.AddCommand(llmCreateCmd)
	llmClientCmd.AddCommand(llmDeleteCmd)
	llmClientCmd.AddCommand(llmListCmd)
	llmClientCmd.AddCommand(llmGetCmd)
	llmClientCmd.AddCommand(llmStopCmd)
	llmClientCmd.AddCommand(llmStartCmd)
	llmClientCmd.AddCommand(llmGenerateCmd)
	vmClientCmd.AddCommand(vmGetCmd)
	vmClientCmd.AddCommand(vmCreateCmd)
	vmClientCmd.AddCommand(vmDeleteCmd)
	vmClientCmd.AddCommand(vmListCmd)
	vmClientCmd.AddCommand(vmStopCmd)
	vmClientCmd.AddCommand(vmStartCmd)
	containerClientCmd.AddCommand(containerCreateCmd)
	containerClientCmd.AddCommand(containerDeleteCmd)
	containerClientCmd.AddCommand(containerListCmd)
	containerClientCmd.AddCommand(containerGetCmd)
	containerClientCmd.AddCommand(containerStopCmd)
	containerClientCmd.AddCommand(containerStartCmd)
	nodeClientCmd.AddCommand(nodeAddCmd)
	nodeClientCmd.AddCommand(nodeDeleteCmd)
	nodeClientCmd.AddCommand(nodeListCmd)
	nodeClientCmd.AddCommand(nodeGetCmd)
	fileClientCmd.AddCommand(fileUploadCmd)
	fileClientCmd.AddCommand(fileDownloadCmd)
	fileClientCmd.AddCommand(fileDeleteCmd)
	fileClientCmd.AddCommand(fileListCmd)
	sshClientCmd.AddCommand(sshNodeCmd)
	sshClientCmd.AddCommand(sshVMCmd)
	k8sClientCmd.AddCommand(k8sCreateCmd)
	k8sClientCmd.AddCommand(k8sDeleteCmd)
	k8sClientCmd.AddCommand(k8sListCmd)
	k8sClientCmd.AddCommand(k8sGetCmd)
	k8sClientCmd.AddCommand(k8sStartCmd)
	k8sClientCmd.AddCommand(k8sStopCmd)
	k8sClientCmd.AddCommand(k8sGetKubeconfigCmd)
	rootCmd.PersistentFlags().StringVar(&clientHost, "host", "master.govno.cloud", "host to connect to")
	rootCmd.PersistentFlags().StringVar(&clientPort, "port", "6969", "port to connect to")
	rootCmd.PersistentFlags().StringVar(&name, "name", "", "name")
	rootCmd.PersistentFlags().StringVar(&image, "image", "", "image")
	rootCmd.PersistentFlags().StringVar(&flavor, "flavor", "", "flavor")
	nodeClientCmd.PersistentFlags().StringVar(&nodehost, "nodehost", "", "node host")
	nodeClientCmd.PersistentFlags().StringVar(&nodeport, "nodeport", "", "node port")
	sshClientCmd.PersistentFlags().StringVar(&user, "user", "root", "user to ssh as")
	sshClientCmd.PersistentFlags().StringVar(&key, "key", "", "ssh key")
	containerClientCmd.PersistentFlags().StringVar(&id, "id", "", "container id")
	fileClientCmd.PersistentFlags().StringVar(&src, "src", "", "file source")
	llmClientCmd.PersistentFlags().StringVar(&model, "model", "", "model")
	llmClientCmd.PersistentFlags().StringVar(&input, "input", "", "input")
}

func main() {
	Execute()
}
