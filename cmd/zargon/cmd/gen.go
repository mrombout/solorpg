package cmd

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Manages your installed generators",
	Long:  "Manages your installed generators.",
}

func init() {
	genCmd.AddCommand(genListCmd)
	genCmd.AddCommand(genInstallCmd)
	genCmd.AddCommand(genUninstallCmd)
}
