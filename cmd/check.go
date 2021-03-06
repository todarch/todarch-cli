package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/todarch/todarch-cli/tclient"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Health Check",
	Long:  "Can be used for debugging the application",
	Run: func(cmd *cobra.Command, args []string) {
		if tclient.IsTdUp() {
			fmt.Println("Todarch Td service is up and running")
		} else {
			fmt.Println("Todarch Td service is not up")
		}
		if tclient.IsUmUp() {
			fmt.Println("Todarch Um service is up and running")
		} else {
			fmt.Println("Todarch Um service is not up")
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
