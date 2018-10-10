package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "todarch cli v0.1.x"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of of Todarch",
	Long:  "Todarch's version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
