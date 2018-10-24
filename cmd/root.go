package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/todarch/todarch-cli/cmd/todo"
	"github.com/todarch/todarch-cli/consts"
	"github.com/todarch/todarch-cli/util"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todarch",
	Short: "Todarch is a todo archive manager",
	Long:  "Collect your todos, and get them done whenever you want.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root command")
	},
}

func init() {

	// set config defaults
	viper.SetDefault(consts.TempDir, os.TempDir())

	rootCmd.PersistentFlags().BoolP(consts.VERBOSE, "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP(consts.DEBUG, "d", false, "debug application")
	viper.BindPFlag(consts.VERBOSE, rootCmd.PersistentFlags().Lookup(consts.VERBOSE))
	viper.BindPFlag(consts.DEBUG, rootCmd.PersistentFlags().Lookup(consts.DEBUG))

	rootCmd.AddCommand(
		todo.NewTodoCommand(),
	)

	preCheck()
	loadConfig()
}

func preCheck() {
	util.BeSureTodarchWorkspaceExists()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadConfig() {
	viper.AddConfigPath(util.GetTodarchWorkspace())
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	util.Debug("Using config: " + viper.ConfigFileUsed())
}
