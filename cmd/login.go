package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
  Use: "login",
  Short: "Login to application",
  Long: "You need to login to operate on your todos",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Going to login...")
  },
}

func init() {
  rootCmd.AddCommand(loginCmd)
}
