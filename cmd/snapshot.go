package cmd

import (
	"github.com/spf13/cobra"
	"main.go/internal"
	"main.go/pkg/exec"
)

func init() {
	rootCmd.AddCommand(snapshotCmd)
	snapshotCmd.Flags().StringP("algo", "a", "", "指定的算法名称")
}

var AValue string

var snapshotCmd = &cobra.Command{
	Use: "snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		internal.SnapshotAValue, _ = cmd.Flags().GetString("algo")
		if internal.SnapshotAValue == "" {
			internal.SnapshotAValue = "sha256"
		}
		exec.Snapshot()
	},
}
