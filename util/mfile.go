package util

import (
	"os"
	"os/user"
)

const sep = string(os.PathSeparator)

func BeSureTodarchWorkspaceExists() {
	workspace := GetTodarchWorkspace()
	if _, err := os.Stat(workspace); os.IsNotExist(err) {
		if err := os.Mkdir(workspace, os.ModePerm); err != nil {
			SayLastWords("Could not create workspace: " + err.Error())
		}
	}
}

func GetTodarchWorkspace() string {
	current, _ := user.Current()
	filename := current.HomeDir + sep + ".todarch"
	Debug("Todarch workspace directory: " + filename)
	return filename
}

func GetTokenStoreFile() string {
	return GetTodarchWorkspace() + sep + "tokenstore"
}
