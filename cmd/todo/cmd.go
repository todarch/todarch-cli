package todo

import (
	"github.com/spf13/cobra"
)

func NewTodoCommand() *cobra.Command {
	todoCmd := &cobra.Command{
		Use:   "todo",
		Short: "Manage todos",
		Long:  "Shows the details of a todo item",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	todoCmd.AddCommand(
		newRemoveCommand(),
		newListCommand(),
		newInspectCommand(),
		newDoneCommand(),
	)

	return todoCmd

}
