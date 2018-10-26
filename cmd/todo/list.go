package todo

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/todarch/todarch-cli/tclient"
	"github.com/todarch/todarch-cli/util"
	"os"
)

var longFormat bool
var allStatus bool
var rsql string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List todos",
		Run: func(cmd *cobra.Command, args []string) {
			listTodos()
		},
	}

	cmd.Flags().BoolVarP(&longFormat, "", "l", false, "use a long listing format")
	cmd.Flags().BoolVarP(&allStatus, "all", "a", false, "show todos with any status")
	cmd.Flags().StringVarP(&rsql, "rsql", "", "", "filter todos using rsql query")
	return cmd
}

func listTodos() {
	var todos []tclient.TodoItem
	if rsql != "" {
		util.Debug(rsql)
		todos = tclient.GetTodosByRsqlQuery(rsql)
	} else {
		if allStatus {
			todos = tclient.CurrentUserTodos()
		} else {
			todos = tclient.GetTodosByRsqlQuery("todoStatus==INITIAL")
		}
	}

	if len(todos) == 0 {
		fmt.Println("You do not have any todos yet.")
	} else {
		getHeader := func() []string {
			if longFormat {
				return []string{"Id", "Title", "Priority", "Status", "InMin"}
			} else {
				return []string{"Id", "Title", "Priority"}
			}
		}

		getRow := func(todo tclient.TodoItem) []string {
			if longFormat {
				return []string{
					fmt.Sprint(todo.Id),
					todo.Title[:util.Min(len(todo.Title), tablewriter.MAX_ROW_WIDTH)],
					fmt.Sprint(todo.Priority),
					todo.Status,
					fmt.Sprint(todo.TimeNeededInMin)}
			} else {
				return []string{
					fmt.Sprint(todo.Id),
					todo.Title[:util.Min(len(todo.Title), tablewriter.MAX_ROW_WIDTH)],
					fmt.Sprint(todo.Priority)}
			}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(getHeader())
		for _, todo := range todos {
			table.Append(getRow(todo))
		}
		table.Render()
	}
}
