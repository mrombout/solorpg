package cmd

import (
	"github.com/spf13/cobra"
)

var genUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls an installed generator",
	Long:  "Uninstalls an installed generator.",
}
