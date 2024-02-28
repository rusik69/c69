/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/spf13/cobra"
)

var clientHost string
var clientPort string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "govnocloud client",
	Long:  `govnocloud client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: ssh|vm|node|container|file [command] [flags]")
	},
}

// sshClientCmd represents the ssh commands
var sshClientCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh to vm or node",
	Long:  `ssh to vm or node`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: ssh [vm|node]")
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
		cmd.PersistentFlags().String("node", "", "node to ssh to")
		node := cmd.PersistentFlags().Lookup("node")
		if node == nil {
			panic("node is required")
		}
		cmd.PersistentFlags().String("user", "", "user to ssh as")
		var userString string
		user := cmd.PersistentFlags().Lookup("user")
		if user == nil {
			userString = "root"
		}
		err := client.SSHNode(clientHost, clientPort, node.Value.String(), userString)
		if err != nil {
			panic(err)
		}
	},
}

// sshVMCmd represents the ssh vm command
var sshVMCmd = &cobra.Command{
	Use:   "vm",
	Short: "ssh to vm",
	Long:  `ssh to vm`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("vm", "", "vm to ssh to")
		vm := cmd.PersistentFlags().Lookup("vm")
		if vm == nil {
			panic("vm is required")
		}
		cmd.PersistentFlags().String("user", "", "user to ssh as")
		user := cmd.PersistentFlags().Lookup("user")
		var userString string
		if user == nil {
			userString = "root"
		}
		err := client.SSHVM(clientHost, clientPort, vm.Value.String(), userString)
		if err != nil {
			panic(err)
		}
	},
}

// nodeAddCmd represents the node add command
var nodeAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add node",
	Long:  `add node`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("nodehost", "", "node host")
		cmd.PersistentFlags().String("nodeport", "", "node port")
		cmd.PersistentFlags().String("name", "", "node name")
		nodehost := cmd.PersistentFlags().Lookup("nodehost")
		if nodehost == nil {
			panic("nodehost is required")
		}
		nodeport := cmd.PersistentFlags().Lookup("nodeport")
		if nodeport == nil {
			panic("nodeport is required")
		}
		name := cmd.PersistentFlags().Lookup("name")
		if name == nil {
			panic("name is required")
		}
		err := client.AddNode(clientHost, clientPort, name.Value.String(), nodehost.Value.String(), nodeport.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// nodeDeleteCmd represents the node delete command
var nodeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete node",
	Long:  `delete node`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("name", "", "node name")
		name := cmd.PersistentFlags().Lookup("name")
		if name == nil {
			panic("name is required")
		}
		err := client.DeleteNode(clientHost, clientPort, name.Value.String())
		if err != nil {
			panic(err)
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
			panic(err)
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
		cmd.PersistentFlags().String("name", "", "node name")
		name := cmd.PersistentFlags().Lookup("name")
		if name == nil {
			panic("name is required")
		}
		node, err := client.GetNode(clientHost, clientPort, name.Value.String())
		if err != nil {
			panic(err)
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
		cmd.PersistentFlags().String("name", "", "vm name")
		cmd.PersistentFlags().String("image", "", "vm image")
		cmd.PersistentFlags().String("flavor", "", "vm flavor")
		vmName := cmd.PersistentFlags().Lookup("name")
		if vmName == nil {
			panic("vm name is required")
		}
		vmImage := cmd.PersistentFlags().Lookup("image")
		if vmImage == nil {
			panic("vm image is required")
		}
		vmFlavor := cmd.PersistentFlags().Lookup("flavor")
		if vmFlavor == nil {
			panic("vm flavor is required")
		}
		id, err := client.CreateVM(clientHost, clientPort, vmName.Value.String(), vmImage.Value.String(), vmFlavor.Value.String())
		if err != nil {
			panic(err)
		}
		fmt.Println("VM created with id " + fmt.Sprint(id))
	},
}

// vmDeleteCmd represents the vm delete command
var vmDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete vm",
	Long:  `delete vm`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("name", "", "vm name")
		name := cmd.PersistentFlags().Lookup("name")
		if name == nil {
			panic("name is required")
		}
		err := client.DeleteVM(clientHost, clientPort, name.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

var vmGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get vm",
	Long:  `get vm`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("name", "", "vm name")
		name := cmd.PersistentFlags().Lookup("name")
		if name == nil {
			panic("name is required")
		}
		vm, err := client.GetVM(clientHost, clientPort, name.Value.String())
		if err != nil {
			panic(err)
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
			panic(err)
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
		cmd.PersistentFlags().String("name", "", "vm name")
		name := cmd.PersistentFlags().Lookup("name")
		if name == nil {
			panic("name is required")
		}
		err := client.StopVM(clientHost, clientPort, name.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// vmStartCmd represents the vm start command
var vmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start vm",
	Long:  `start vm`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("name", "", "vm name")
		name := cmd.PersistentFlags().Lookup("name")
		if name == nil {
			panic("name is required")
		}
		err := client.StartVM(clientHost, clientPort, name.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// containerCreateCmd represents the container create command
var containerCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create container",
	Long:  `create container`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("name", "", "container name")
		cmd.PersistentFlags().String("image", "", "container image")
		containerName := cmd.PersistentFlags().Lookup("name")
		if containerName == nil {
			panic("container name is required")
		}
		containerImage := cmd.PersistentFlags().Lookup("image")
		if containerImage == nil {
			panic("container image is required")
		}
		id, err := client.CreateContainer(clientHost, clientPort, containerName.Value.String(), containerImage.Value.String())
		if err != nil {
			panic(err)
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
		cmd.PersistentFlags().String("id", "", "container id")
		id := cmd.PersistentFlags().Lookup("id")
		if id == nil {
			panic("id is required")
		}
		err := client.DeleteContainer(clientHost, clientPort, id.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// containerGetCmd represents the container get command
var containerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get container",
	Long:  `get container`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("id", "", "container id")
		id := cmd.PersistentFlags().Lookup("id")
		if id == nil {
			panic("id is required")
		}
		container, err := client.GetContainer(clientHost, clientPort, id.Value.String())
		if err != nil {
			panic(err)
		}
		fmt.Println(
			"ID: " + fmt.Sprint(container.ID) + "\n" +
				"Name: " + container.Name + "\n" +
				"IP: " + container.IP + "\n" +
				"Host: " + container.Host + "\n" +
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
			panic(err)
		}
		fmt.Printf("| %-10s | %-10s | %-16s | %-5s | %-7s | %-20s |\n", "ID", "NAME", "IP", "Host", "Status", "Image")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		for _, container := range containers {
			fmt.Printf("| %-10s | %-10s | %-16s | %-5s | %-7s | %-20s |\n", container.ID, container.Name, container.IP, container.Host, container.State, container.Image)
		}
	},
}

// containerStopCmd represents the container stop command
var containerStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop container",
	Long:  `stop container`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("id", "", "container id")
		containerID := cmd.PersistentFlags().Lookup("id")
		if containerID == nil {
			panic("container id is required")
		}
		err := client.StopContainer(clientHost, clientPort, containerID.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// containerStartCmd represents the container start command
var containerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start container",
	Long:  `start container`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("id", "", "container id")
		containerID := cmd.PersistentFlags().Lookup("id")
		if containerID == nil {
			panic("container id is required")
		}
		err := client.StartContainer(clientHost, clientPort, containerID.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// fileUploadCmd represents the file upload command
var fileUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file",
	Long:  `upload file`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("src", "", "file source")
		fileSource := cmd.PersistentFlags().Lookup("src")
		if fileSource == nil {
			panic("file source is required")
		}
		err := client.UploadFile(clientHost, clientPort, fileSource.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// fileDownloadCmd represents the file download command
var fileDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download file",
	Long:  `download file`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("source", "", "file source")
		fileSource := cmd.PersistentFlags().Lookup("source")
		if fileSource == nil {
			panic("file source is required")
		}
		err := client.DownloadFile(clientHost, clientPort, fileSource.Value.String())
		if err != nil {
			panic(err)
		}
	},
}

// fileDeleteCmd represents the file delete command
var fileDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete file",
	Long:  `delete file`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().String("name", "", "file name")
		fileName := cmd.PersistentFlags().Lookup("name")
		if fileName == nil {
			panic("name is required")
		}
		err := client.DeleteFile(clientHost, clientPort, fileName.Value.String())
		if err != nil {
			panic(err)
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
			panic(err)
		}
		fmt.Println(files)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(sshClientCmd)
	clientCmd.AddCommand(vmClientCmd)
	clientCmd.AddCommand(nodeClientCmd)
	clientCmd.AddCommand(containerClientCmd)
	clientCmd.AddCommand(fileClientCmd)
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	hostVal := clientCmd.PersistentFlags().String("host", "127.0.0.1", "host to connect to")
	portVal := clientCmd.PersistentFlags().String("port", "7070", "port to connect to")
	clientHost = *hostVal
	clientPort = *portVal
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
