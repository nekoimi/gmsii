package cmd

import (
	"fmt"
	"github.com/nekoimi/gmsii/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Example:gmsii version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gmsii " + config.Version)
	},
}
