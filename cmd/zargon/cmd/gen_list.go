package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var genListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available generators on the default generator marketplace",
	Long:  "Lists all available generators on the default generator marketplace.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func init() {
	genListCmd.Flags().Bool("installed", false, "list only installed generators")
}
