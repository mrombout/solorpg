package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initializes a directory to be used as a Zargon story",
	Long:  "Initializes a directory to be used as a Zargon story.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	initCmd.Flags().IntP("seed", "s", 0, "seed used to randomize generators")
}
