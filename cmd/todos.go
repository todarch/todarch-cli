package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/todarch/todarch-cli/tclient"
	"github.com/todarch/todarch-cli/util"
	"os"
)

var todosCmd = &cobra.Command{
	Use:   "todos",
	Short: "Operates on todo items",
	Long:  "Create, delete, request your todos",
	Run: func(cmd *cobra.Command, args []string) {
		todos := tclient.CurrentUserTodos()
		if len(todos) == 0 {
			fmt.Println("You do not have any todos yet.")
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Id", "Title", "Priority"})
			for _, todo := range todos {
				table.Append([]string{
					fmt.Sprint(todo.Id),
					todo.Title[:util.Min(len(todo.Title), tablewriter.MAX_ROW_WIDTH)],
					fmt.Sprint(todo.Priority)})
			}
			table.Render()
		}
	},
}

func init() {
	getCmd.AddCommand(todosCmd)
}
