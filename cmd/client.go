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
		fmt.Println(args)
		if len(args) == 0 {
			panic("node is required")
		}
		node := args[0]
		userPtr := cmd.PersistentFlags().String("user", "root", "user to ssh as")
		if *userPtr == "" {
			panic("user is required")
		}
		err := client.SSHNode(clientHost, clientPort, node, *userPtr)
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
		vmPtr := cmd.PersistentFlags().String("vm", "", "vm to ssh to")
		if *vmPtr == "" {
			panic("vm is required")
		}
		userPtr := cmd.PersistentFlags().String("user", "root", "user to ssh as")
		if *userPtr == "" {
			panic("user is required")
		}
		err := client.SSHVM(clientHost, clientPort, *vmPtr, *userPtr)
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
		hostPtr := cmd.PersistentFlags().String("nodehost", "", "node host")
		if *hostPtr == "" {
			panic("nodehost is required")
		}
		portPtr := cmd.PersistentFlags().String("nodeport", "", "node port")
		if *portPtr == "" {
			panic("nodeport is required")
		}
		namePtr := cmd.PersistentFlags().String("name", "", "node name")
		if *namePtr == "" {
			panic("name is required")
		}
		err := client.AddNode(clientHost, clientPort, *namePtr, *portPtr, *namePtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "node name")
		if *namePtr == "" {
			panic("name is required")
		}
		err := client.DeleteNode(clientHost, clientPort, *namePtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "node name")
		if *namePtr == "" {
			panic("name is required")
		}
		node, err := client.GetNode(clientHost, clientPort, *namePtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "vm name")
		if *namePtr == "" {
			panic("name is required")
		}
		imagePtr := cmd.PersistentFlags().String("image", "", "vm image")
		if *imagePtr == "" {
			panic("image is required")
		}
		flavorPtr := cmd.PersistentFlags().String("flavor", "", "vm flavor")
		if *flavorPtr == "" {
			panic("flavor is required")
		}
		id, err := client.CreateVM(clientHost, clientPort, *namePtr, *imagePtr, *flavorPtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "vm name")
		if *namePtr == "" {
			panic("name is required")
		}
		err := client.DeleteVM(clientHost, clientPort, *namePtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "vm name")
		if *namePtr == "" {
			panic("name is required")
		}
		vm, err := client.GetVM(clientHost, clientPort, *namePtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "vm name")
		if *namePtr == "" {
			panic("name is required")
		}
		err := client.StopVM(clientHost, clientPort, *namePtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "vm name")
		if *namePtr == "" {
			panic("name is required")
		}
		err := client.StartVM(clientHost, clientPort, *namePtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "container name")
		if *namePtr == "" {
			panic("name is required")
		}
		imagePtr := cmd.PersistentFlags().String("image", "", "container image")
		if *imagePtr == "" {
			panic("image is required")
		}
		id, err := client.CreateContainer(clientHost, clientPort, *namePtr, *imagePtr)
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
		idPtr := cmd.PersistentFlags().String("id", "", "container id")
		if *idPtr == "" {
			panic("id is required")
		}
		err := client.DeleteContainer(clientHost, clientPort, *idPtr)
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
		idPtr := cmd.PersistentFlags().String("id", "", "container id")
		if *idPtr == "" {
			panic("id is required")
		}
		container, err := client.GetContainer(clientHost, clientPort, *idPtr)
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
		idPtr := cmd.PersistentFlags().String("id", "", "container id")
		if *idPtr == "" {
			panic("id is required")
		}
		err := client.StopContainer(clientHost, clientPort, *idPtr)
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
		idPtr := cmd.PersistentFlags().String("id", "", "container id")
		err := client.StartContainer(clientHost, clientPort, *idPtr)
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
		srcPtr := cmd.PersistentFlags().String("src", "", "file source")
		if *srcPtr == "" {
			panic("src is required")
		}
		err := client.UploadFile(clientHost, clientPort, *srcPtr)
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
		srcPtr := cmd.PersistentFlags().String("source", "", "file source")
		if *srcPtr == "" {
			panic("source is required")
		}
		err := client.DownloadFile(clientHost, clientPort, *srcPtr)
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
		namePtr := cmd.PersistentFlags().String("name", "", "file name")
		if *namePtr == "" {
			panic("name is required")
		}
		err := client.DeleteFile(clientHost, clientPort, *namePtr)
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
	clientCmd.PersistentFlags().StringVar(&clientHost, "host", "127.0.0.1", "host to connect to")
	clientCmd.PersistentFlags().StringVar(&clientPort, "port", "7070", "port to connect to")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
