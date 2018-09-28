package cmd

import (
	"github.com/spf13/cobra"
	"github.com/todarch/todarch-cli/tclient"
)

func init() {
	var todoCmd = &cobra.Command{
		Use:   "todo",
		Short: "Operates on todo items",
		Long:  "Create, delete, request your todos",
		Run: func(cmd *cobra.Command, args []string) {
			tclient.CurrentUserTodos()
		},
	}
	rootCmd.AddCommand(todoCmd)
}
