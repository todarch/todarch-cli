package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/todarch/todarch-cli/consts"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todarch",
	Short: "Todarch is a todo achieve manager",
	Long:  "Collect your todos, and get them done whenever you want.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root command")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP(consts.VERBOSE, "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP(consts.DEBUG, "d", false, "debug application")
	viper.BindPFlag(consts.VERBOSE, rootCmd.PersistentFlags().Lookup(consts.VERBOSE))
	viper.BindPFlag(consts.DEBUG, rootCmd.PersistentFlags().Lookup(consts.DEBUG))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
