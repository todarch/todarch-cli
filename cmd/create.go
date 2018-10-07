package cmd

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/todarch/todarch-cli/consts"
	"github.com/todarch/todarch-cli/tclient"
	"github.com/todarch/todarch-cli/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
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
				// show a ready-to-fill yaml template to user their editor
				// when they close the close, validate content
				// open the editor again, with errors at the top in comment format
				// let them fix the errors, or use aborting directive at the top to exit
				filename := uniqueFullFilePath()
				createInitialTemplate(filename)
				editInEditor(filename)
				valid := false
				for !valid {
					todoReq = getFromFile(filename)
					allErrors := validateTodoReq(todoReq)
					if len(allErrors) == 0 {
						valid = true
					} else {
						insertYmlTemplate(filename, todoReq, allErrors)
						editInEditor(filename)
					}
				}
			}
			tclient.NewTodo(todoReq)
		},
	}
	createCmd.Flags().StringVarP(&file, "file", "f", "", "file defining new todo")
	createCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "define a new todo interactively")
	createCmd.Flags().BoolVarP(&editor, "editor", "e", true, "define a new todo in editor")
	rootCmd.AddCommand(createCmd)
}

func validateTodoReq(req tclient.TodoCreationReq) (allErrors []error) {
	if len(strings.TrimSpace(req.Title)) == 0 {
		allErrors = append(allErrors, errors.New("Title is required"))
	}

	return allErrors
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

	//TODO:selimssevgi: may not be a good idea to stop the flow in the middle somewhere
	readableContent := string(fileContent)
	if strings.HasPrefix(readableContent, "#EXIT") {
		util.SayLastWords("File content starts with abort directive.")
	}

	var todoReq tclient.TodoCreationReq
	err = yaml.Unmarshal(fileContent, &todoReq)
	if err != nil {
		fmt.Println("Could not parse file content:", err)
	}
	util.Debug(todoReq)
	return todoReq
}

func editInEditor(filename string) {
	userEditor := os.Getenv("EDITOR")
	if userEditor == "" {
		util.SayLastWords("Set $EDITOR env variable to use this functionality")
	}
	cmd := exec.Command(userEditor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		util.SayLastWords(err.Error())
	}
}

func uniqueFullFilePath() string {
	t := time.Now()
	uniqueFilePath := os.TempDir() + "/todo-" + t.Format("20060102150405") + ".yml"
	util.Debug("Unique full file path: " + uniqueFilePath)
	return uniqueFilePath
}

func createInitialTemplate(fullPathFileName string) {
	var emptyTodo tclient.TodoCreationReq
	insertYmlTemplate(fullPathFileName, emptyTodo, nil)
}

func insertYmlTemplate(fullPathFileName string, todoReq tclient.TodoCreationReq, allErrors []error) {
	out, err := yaml.Marshal(todoReq)
	if err != nil {
		util.SayLastWords(err.Error())
	}
	errorOutput := formatErrors(allErrors)
	fileContent := errorOutput + string(out)
	err = ioutil.WriteFile(fullPathFileName, []byte(fileContent), 0644)
	if err != nil {
		util.SayLastWords(err.Error())
	}
}

// Formats given error list into something that can be included
// in a yaml file as comment. Informs users about the erroneous parts.
func formatErrors(allErrors []error) string {
	if len(allErrors) == 0 {
		return ""
	}
	errorBegin := "## todarch:errors"
	errorOutput := "" + errorBegin
	for _, err := range allErrors {
		errorOutput = errorOutput + "\n # - " + err.Error()
	}
	return errorOutput + "\n\n"
}
