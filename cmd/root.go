package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todarch",
	Short: "Todarch is a todo archieve manager",
	Long:  "Collect your todos, and get them done whenever you want.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
