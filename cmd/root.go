package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"main.go/internal/version"
)

var rootCmd = &cobra.Command{
	Use:     "vec",
	Short:   "Vec is short for verdict,is a CLI tool used to verify file integrity.",
	Version: version.Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
