/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

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

// vmCmd represents the vm commands
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "vm commands",
	Long:  `vm commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: vm [create|delete|list|get]")
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
		fmt.Printf("| %-10s | %-10s | %-10s |\n", "NAME", "HOST", "PORT")
		fmt.Println("------------------------------------------------")
		for _, node := range nodes {
			fmt.Printf("| %-10s | %-10s | %-10s |\n", node.Name, node.Host, node.Port)
		}
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
		idString := cmd.PersistentFlags().Lookup("id").Value.String()
		if idString == "" {
			panic("id is required")
		}
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		err = client.DeleteVM(host, port, id)
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
		idString := cmd.PersistentFlags().Lookup("id").Value.String()
		if idString == "" {
			panic("id is required")
		}
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		vm, err := client.GetVM(host, port, id)
		if err != nil {
			panic(err)
		}
		fmt.Println(
			"ID: " + fmt.Sprint(vm.ID) + "\n" +
				"Name: " + vm.Name + "\n" +
				"IP: " + vm.IP + "\n" +
				"Host: " + vm.Host + "\n" +
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
		fmt.Printf("| %-10s | %-10s | %-16s | %-5s | %-7s | %-20s | %-5s | %-10s |\n", "ID", "NAME", "IP", "Host", "Status", "Image", "Flavor", "Volumes")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		for _, vm := range vms {
			fmt.Printf("| %-10d | %-10s | %-16s | %-5s | %-7s | %-20s | %-5s | %-10s |\n", vm.ID, vm.Name, vm.IP, vm.Host, vm.State, vm.Image, vm.Flavor, vm.Volumes)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(vmCmd)
	clientCmd.AddCommand(nodeClientCmd)
	vmCmd.AddCommand(vmGetCmd)
	vmCmd.AddCommand(vmCreateCmd)
	vmCmd.AddCommand(vmDeleteCmd)
	vmCmd.AddCommand(vmListCmd)
	nodeClientCmd.AddCommand(nodeAddCmd)
	nodeClientCmd.AddCommand(nodeDeleteCmd)
	nodeClientCmd.AddCommand(nodeListCmd)
	nodeClientCmd.AddCommand(nodeGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	clientCmd.PersistentFlags().String("host", "localhost", "host to connect to")
	clientCmd.PersistentFlags().String("port", "6969", "port to connect to")
	vmCmd.PersistentFlags().String("name", "", "name of the vm")
	vmCmd.PersistentFlags().String("image", "", "image of the vm")
	vmCmd.PersistentFlags().String("flavor", "", "flavor of the vm")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
