/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/maniizu3110/go-codegen/start"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create template codegen.yaml",
	Long:  `go-codegen will auto-generate according to the configuration file. init will automatically create the recommended codegen.yaml for this purpose.`,
	Run: func(cmd *cobra.Command, args []string) {
		start.CreateDefatultYaml()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
