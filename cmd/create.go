package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/todarch/todarch-cli/consts"
	"github.com/todarch/todarch-cli/tclient"
	"github.com/todarch/todarch-cli/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var file string
var interactive bool
var editor bool

func init() {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "create todo",
		Long:  "create new todo",
		Run: func(cmd *cobra.Command, args []string) {
			var todoReq tclient.TodoCreationReq
			if file != "" {
				todoReq = getFromFile(file)
			} else if interactive {
				todoReq = getInteractively()
				saveToTemp(todoReq)
			} else if editor {
				util.SayLastWords("Gonna implement editor option later")
			}
			tclient.NewTodo(todoReq)
		},
	}
	createCmd.Flags().StringVarP(&file, "file", "f", "", "file defining new todo")
	createCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "define a new todo interactively")
	createCmd.Flags().BoolVarP(&editor, "editor", "e", true, "define a new todo in editor")
	rootCmd.AddCommand(createCmd)
}
func saveToTemp(todoCreationReq tclient.TodoCreationReq) {
	out, err := yaml.Marshal(todoCreationReq)
	if err != nil {
		util.Debug("Could not convert interactive todo to yaml output: " + err.Error())
		return
	}

	fileName := strings.Replace(todoCreationReq.Title, " ", "_", -1)
	fileFullPath := viper.GetString(consts.TempDir) + "/" + fileName + ".yml"
	err = ioutil.WriteFile(fileFullPath, out, 0644)
	if err != nil {
		util.Debug("Could not write temp yaml output: " + err.Error())
		return
	}
	fmt.Println("Wrote todo req to ", fileFullPath)
}

func getInteractively() tclient.TodoCreationReq {
	var todoReq tclient.TodoCreationReq
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Title: ")
	todoReq.Title = getText(reader)

	fmt.Print("Description: ")
	todoReq.Description = getText(reader)

	fmt.Print("Priority: ")
	priorityAsStr := getText(reader)
	todoReq.Priority = cast.ToInt(priorityAsStr)

	return todoReq
}

func getText(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return str[:len(str)-1]
}

func getFromFile(filename string) tclient.TodoCreationReq {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Could not read file", filename, err)
	}

	var todoReq tclient.TodoCreationReq
	err = yaml.Unmarshal(fileContent, &todoReq)
	if err != nil {
		fmt.Println("Could not parse file content:", err)
	}
	util.Debug(todoReq)
	return todoReq
}
