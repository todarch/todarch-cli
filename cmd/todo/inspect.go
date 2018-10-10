package todo

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/todarch/todarch-cli/tclient"
	"github.com/todarch/todarch-cli/util"
	"gopkg.in/yaml.v2"
)

var outputFormat string

func newInspectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "Display detailed information on a todo item",
		Long:  "Shows the details of a todo item",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				util.SayLastWords("You need to give an id.")
			}

			todoId := cast.ToInt(args[0])
			if todoId == 0 || todoId < 0 {
				util.SayLastWords("Id should be a positive number: " + args[0])
			}

			if len(args) == 2 && args[1] == "done" {
				tclient.GetTodoDone(todoId)
			} else {
				todoDetail := tclient.TodoById(todoId)
				printInOutputFormat(todoDetail, outputFormat)
			}
		},
	}

	cmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "output format")

	return cmd
}

func printInOutputFormat(item tclient.TodoItem, outputFormat string) {
	switch outputFormat {
	case "yml", "yaml":
		printInYml(item)
	case "json":
		printInJson(item)
	default:
		printInJson(item)
	}
}

func printInJson(item tclient.TodoItem) {
	out, err := json.MarshalIndent(item, "", "    ")
	if err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println("Something unexpected happen")
		util.Debug(err)
	}
}

func printInYml(item tclient.TodoItem) {
	out, err := yaml.Marshal(item)
	if err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println("Something unexpected happen")
		util.Debug(err)
	}
}
