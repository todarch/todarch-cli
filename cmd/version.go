package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of of Todarch",
	Long:  "Todarch's version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todarch CLI v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
