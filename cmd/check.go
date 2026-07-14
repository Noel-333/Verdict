package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"main.go/pkg/exec"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "TODO: 什么都还没写呢",
	Run: func(cmd *cobra.Command, args []string) {
		if exec.CheckMeta() {
			exec.CheckData()
		} else {
			color.Red("Meta data error.")
		}
	},
}
