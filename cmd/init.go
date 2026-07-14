package cmd

import (
	"github.com/spf13/cobra"
	"main.go/pkg/exec"
)

func init() {
	rootCmd.AddCommand(initCmd)

}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Run init in the terminal to create configuration files in your current working directory.",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Initialize()
	},
}
