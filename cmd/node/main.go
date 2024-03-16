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
	"github.com/rusik69/govnocloud/pkg/node"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govnocloud",
	Short: "govnocloud is a shitty cloud",
	Long:  `govnocloud is a shitty cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		envInstance, err := node.ParseEnv()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		types.NodeEnvInstance = envInstance
		logrus.Println("Node environment is parsed")
		logrus.Println("Node name is " + types.NodeEnvInstance.Name)
		logrus.Println("Node IP is " + types.NodeEnvInstance.IP)
		logrus.Println("Node port is " + types.NodeEnvInstance.ListenPort)
		logrus.Println("Node libvirt socket is " + types.NodeEnvInstance.LibVirtURI)
		logrus.Println("Node libvirt image dir is " + types.NodeEnvInstance.LibVirtImageDir)
		node.LibvirtConnection, err = node.VMConnect()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		defer node.LibvirtConnection.Close()
		node.DockerConnection, err = node.ContainerConnect()
		defer node.DockerConnection.Close()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		err = node.CreateSSHKey()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		err = node.DownloadImages()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		err = node.InstallAnsible()
		if err != nil {
			logrus.Error(err.Error())
			panic(err)
		}
		node.Serve()
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
