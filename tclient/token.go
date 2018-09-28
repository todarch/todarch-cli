package tclient

import (
	"github.com/todarch/todarch-cli/util"
	"io/ioutil"
	"os"
)

func tokenFilename() string {
	todarchHome := os.Getenv("HOME")
	filename := todarchHome + "/.todarch/tokenstore"
	return filename
}

func saveToken(token string) {
	if token == "" {
		util.Log("Could not find the authorization token from headers.")
		os.Exit(-2)
	}
	err := ioutil.WriteFile(tokenFilename(), []byte(token), 0644)
	util.PanicOnError(err)
	util.Debug("Wrote authorization token to token store.")
}

func readToken() string {
	token, err := ioutil.ReadFile(tokenFilename())
	if err != nil {
		util.Debug("Could not read token, returning empty string")
		return ""
	}
	return string(token)
}

// clearToken clears the saved token.
// It could be used for an expired token, or any other reason.
func clearToken() {
	err := ioutil.WriteFile(tokenFilename(), []byte(""), 0644)
	util.PanicOnError(err)
	util.Debug("Cleared token")
}
