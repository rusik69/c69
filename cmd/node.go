/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rusik69/govnocloud/pkg/node/env"
	"github.com/rusik69/govnocloud/pkg/node/server"
	"github.com/rusik69/govnocloud/pkg/node/vm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Start govnocloud node",
	Long:  `Start govnocloud node.`,
	Run: func(cmd *cobra.Command, args []string) {
		envInstance, err := env.Parse()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		env.NodeEnvInstance = envInstance
		vm.LibvirtConnection, err = vm.Connect()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		defer vm.LibvirtConnection.Close()
		err = vm.DownloadImages()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
