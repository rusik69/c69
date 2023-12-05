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

// vm represents the vm commands
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "vm commands",
	Long:  `vm commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage: vm [create|delete|list]")
	},
}

// vmCreateCmd represents the vm create command
var vmCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create vm",
	Long:  `create vm`,
	Run: func(cmd *cobra.Command, args []string) {
		err := client.CreateVM(cmd)
		if err != nil {
			panic(err)
		}
	},
}

// vmDeleteCmd represents the vm delete command
var vmDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete vm",
	Long:  `delete vm`,
	Run: func(cmd *cobra.Command, args []string) {
		err := client.DeleteVM(cmd)
		if err != nil {
			panic(err)
		}
	},
}

// vmListCmd represents the vm list command
var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "list vm",
	Long:  `list vm`,
	Run: func(cmd *cobra.Command, args []string) {
		vms, err := client.ListVMs(cmd)
		if err != nil {
			panic(err)
		}
		fmt.Printf("| %-10s | %-10s | %-16s | %-5s | %-7s | %-20s | %-5s | %-10s |\n", "ID", "NAME", "IP", "Host", "Status", "Image", "Flavor", "Volumes")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		for _, vm := range vms {
			fmt.Println("| %-10s | %-10s | %-16s | %-5s | %-7s | %-20s | %-5s | %-10s |\n", vm.ID, vm.Name, vm.IP, vm.Host, vm.State, vm.Image, vm.Flavor, vm.Volumes)
		}
	},
}

func init() {

	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(vmCmd)
	vmCmd.AddCommand(vmCreateCmd)
	vmCmd.AddCommand(vmDeleteCmd)
	vmCmd.AddCommand(vmListCmd)

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
