package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/todarch/tclient"
	"os"
	"strings"
)

func init() {
	var username string
	var password string

	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to application",
		Long:  "You need to login to operate on your todos",
		Run: func(cmd *cobra.Command, args []string) {
			if username == "" || password == "" {
				username, password = credentials()
			}
			tclient.Authenticate(username, password)
		},
	}
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "your login username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "your login password")
	rootCmd.AddCommand(loginCmd)
}

func credentials() (username string, password string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	rawUsername, _ := reader.ReadString('\n')
	//TODO:selimssevgi: do not show password when typing
	fmt.Print("Enter password: ")
	rawPassword, _ := reader.ReadString('\n')

	username = strings.TrimSpace(rawUsername)
	password = strings.TrimSpace(rawPassword)
	return
}
