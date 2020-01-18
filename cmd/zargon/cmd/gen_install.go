package cmd

import (
	"github.com/spf13/cobra"
)

var genInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs an available generator on the default generator marketplace",
	Long:  "Installs an available generator on the default generator marketplace.",
}
