package todo

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/todarch/todarch-cli/tclient"
	"github.com/todarch/todarch-cli/util"
)

func newRemoveCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a todo",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				util.SayLastWords("You need to give an id.")
			}

			todoId := cast.ToInt(args[0])
			if todoId == 0 || todoId < 0 {
				util.SayLastWords("Id should be a positive number: " + args[0])
			}

			tclient.DeleteTodoById(todoId)
		},
	}

	return cmd
}
