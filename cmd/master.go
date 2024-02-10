/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rusik69/govnocloud/pkg/master"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "Start govnocloud master",
	Long:  `Start govnocloud master.`,
	Run: func(cmd *cobra.Command, args []string) {
		envInstance, err := master.ParseEnv()
		if err != nil {
			panic(err)
		}
		types.MasterEnvInstance = envInstance
		logrus.Println("Master environment is parsed")
		logrus.Println("ETCD host is " + types.MasterEnvInstance.ETCDHost)
		logrus.Println("ETCD port is " + types.MasterEnvInstance.ETCDPort)
		logrus.Println("ETCD user is " + types.MasterEnvInstance.ETCDUser)
		logrus.Println("ETCD pass is " + types.MasterEnvInstance.ETCDPass)
		logrus.Println("Listen port is " + types.MasterEnvInstance.ListenPort)
		master.ETCDClient, err = master.ETCDConnect(types.MasterEnvInstance.ETCDHost,
			types.MasterEnvInstance.ETCDPort, types.MasterEnvInstance.ETCDUser,
			types.MasterEnvInstance.ETCDPass)
		if err != nil {
			panic(err)
		}
		defer master.ETCDClient.Close()
		logrus.Println("ETCD is connected at " + types.MasterEnvInstance.ETCDHost + ":" + types.MasterEnvInstance.ETCDPort)
		master.Serve()
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
