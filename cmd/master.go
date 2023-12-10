/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rusik69/govnocloud/pkg/master/env"
	"github.com/rusik69/govnocloud/pkg/master/server"
	"github.com/spf13/cobra"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "Start govnocloud master",
	Long:  `Start govnocloud master.`,
	Run: func(cmd *cobra.Command, args []string) {
		envInstance, err := env.Parse()
		if err != nil {
			panic(err)
		}
		env.MasterEnvInstance = envInstance
		server.ETCDClient, err = server.ETCDConnect(env.MasterEnvInstance.ETCDHost, env.MasterEnvInstance.ETCDPort, env.MasterEnvInstance.ETCDUser, env.MasterEnvInstance.ETCDPass)
		if err != nil {
			panic(err)
		}
		defer server.ETCDClient.Close()

	},
}

func init() {
	rootCmd.AddCommand(masterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// masterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// masterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
