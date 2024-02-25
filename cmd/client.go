/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "govnocloud client",
	Long:  `govnocloud client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		node := cmd.PersistentFlags().Lookup("node").Value.String()
		if node == "" {
			panic("node is required")
		}
		user := cmd.PersistentFlags().Lookup("user").Value.String()
		if user == "" {
			user = "root"
		}
		err := client.SSHNode(host, port, node, user)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		vm := cmd.PersistentFlags().Lookup("vm").Value.String()
		if vm == "" {
			panic("vm is required")
		}
		user := cmd.PersistentFlags().Lookup("user").Value.String()
		if user == "" {
			user = "root"
		}
		err := client.SSHVM(host, port, vm, user)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		nodehost := cmd.PersistentFlags().Lookup("nodehost").Value.String()
		if host == "" {
			panic("nodehost is required")
		}
		nodeport := cmd.PersistentFlags().Lookup("nodeport").Value.String()
		if port == "" {
			panic("nodeport is required")
		}
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("name is required")
		}
		err := client.AddNode(nodehost, nodeport, host, port, name)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("name is required")
		}
		err := client.DeleteNode(host, port, name)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		nodes, err := client.ListNodes(host, port)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("name is required")
		}
		node, err := client.GetNode(host, port, name)
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
		vmName := cmd.PersistentFlags().Lookup("name").Value.String()
		if vmName == "" {
			panic("vm name is required")
		}
		vmImage := cmd.PersistentFlags().Lookup("image").Value.String()
		if vmImage == "" {
			panic("vm image is required")
		}
		vmFlavor := cmd.PersistentFlags().Lookup("flavor").Value.String()
		if vmFlavor == "" {
			panic("vm flavor is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		id, err := client.CreateVM(host, port, vmName, vmImage, vmFlavor)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("name is required")
		}
		err := client.DeleteVM(host, port, name)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("name is required")
		}
		vm, err := client.GetVM(host, port, name)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		vms, err := client.ListVMs(host, port)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("name is required")
		}
		err := client.StopVM(host, port, name)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("name is required")
		}
		err := client.StartVM(host, port, name)
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
		containerName := cmd.PersistentFlags().Lookup("name").Value.String()
		if containerName == "" {
			panic("container name is required")
		}
		containerImage := cmd.PersistentFlags().Lookup("image").Value.String()
		if containerImage == "" {
			panic("container image is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		id, err := client.CreateContainer(host, port, containerName, containerImage)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		id := cmd.PersistentFlags().Lookup("id").Value.String()
		if id == "" {
			panic("id is required")
		}
		err := client.DeleteContainer(host, port, id)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		id := cmd.PersistentFlags().Lookup("id").Value.String()
		if id == "" {
			panic("id is required")
		}
		container, err := client.GetContainer(host, port, id)
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
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		containers, err := client.ListContainers(host, port)
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
		containerID := cmd.PersistentFlags().Lookup("id").Value.String()
		if containerID == "" {
			panic("container id is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}

		err := client.StopContainer(host, port, containerID)
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
		containerID := cmd.PersistentFlags().Lookup("id").Value.String()
		if containerID == "" {
			panic("container id is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}

		err := client.StartContainer(host, port, containerID)
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
		fileSource := cmd.PersistentFlags().Lookup("src").Value.String()
		if fileSource == "" {
			panic("file source is required")
		}
		fileDestination := cmd.PersistentFlags().Lookup("dest").Value.String()
		if fileDestination == "" {
			panic("file destination is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		err := client.UploadFile(host, port, fileSource)
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
		fileSource := cmd.PersistentFlags().Lookup("source").Value.String()
		if fileSource == "" {
			panic("file source is required")
		}
		fileDestination := cmd.PersistentFlags().Lookup("destination").Value.String()
		if fileDestination == "" {
			panic("file destination is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		err := client.DownloadFile(host, port, fileSource)
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
		fileName := cmd.PersistentFlags().Lookup("name").Value.String()
		if fileName == "" {
			panic("name is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}

		err := client.DeleteFile(host, port, fileName)
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
		name := cmd.PersistentFlags().Lookup("name").Value.String()
		if name == "" {
			panic("file name is required")
		}
		host := cmd.PersistentFlags().Lookup("host").Value.String()
		if host == "" {
			panic("host is required")
		}
		port := cmd.PersistentFlags().Lookup("port").Value.String()
		if port == "" {
			panic("port is required")
		}
		files, err := client.ListFiles(host, port)
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
	clientCmd.PersistentFlags().String("host", "localhost", "host to connect to")
	clientCmd.PersistentFlags().String("port", "6969", "port to connect to")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
