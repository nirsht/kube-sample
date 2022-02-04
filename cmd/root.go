/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kube-sample",
	Short: "A brief description of your application",
	Long: `kube-sample is a CLI for creating and running sample yaml.
This tool can use for development or for lite templating`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().BoolP("run", "r", false, "Also run the yaml file with kubectl")
	createCmd.PersistentFlags().BoolP("apply", "a", false, "Also run the yaml file with kubectl, also if exists on the cluster")
}
