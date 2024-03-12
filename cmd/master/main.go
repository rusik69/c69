/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/simplecloud/pkg/master"
	"github.com/rusik69/simplecloud/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simplecloud",
	Short: "simplecloud is a simple cloud",
	Long:  `simplecloud is a simple cloud`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := strings.Split(f.File, "/")
			return fmt.Sprintf("%s:%d", filename[len(filename)-1], f.Line), ""
		},
	})
	gin.DefaultWriter = logrus.StandardLogger().Writer()
	gin.DefaultErrorWriter = logrus.StandardLogger().Writer()
}

func main() {
	Execute()
}
