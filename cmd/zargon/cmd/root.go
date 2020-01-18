package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zargon",
	Short: "Zargon is a utility to set-up and manage your story.",
	Long:  `Zargon is a utility to set-up and manage your story.`,
}

func init() {
	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(genCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
