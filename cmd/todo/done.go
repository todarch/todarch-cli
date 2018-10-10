package todo

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/todarch/todarch-cli/tclient"
	"github.com/todarch/todarch-cli/util"
)

func newDoneCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "done",
		Short: "Mark a todo item as done",
		Run: func(cmd *cobra.Command, args []string) {
			runDone(args)
		},
	}

	return cmd
}

func runDone(args []string) {
	if len(args) == 0 {
		util.SayLastWords("You need to give an id.")
	}

	todoId := cast.ToInt(args[0])
	if todoId == 0 || todoId < 0 {
		util.SayLastWords("Id should be a positive number: " + args[0])
	}

	tclient.GetTodoDone(todoId)
}
