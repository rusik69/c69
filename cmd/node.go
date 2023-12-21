/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rusik69/govnocloud/pkg/node"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Start govnocloud node",
	Long:  `Start govnocloud node.`,
	Run: func(cmd *cobra.Command, args []string) {
		envInstance, err := node.ParseEnv()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		types.NodeEnvInstance = envInstance
		node.LibvirtConnection, err = node.VMConnect()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		defer node.LibvirtConnection.Close()
		err = node.DownloadVMImages()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		node.Serve()
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
