package main

import "os"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a shitty cloud",
	Long:  `deploy a shitty cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
