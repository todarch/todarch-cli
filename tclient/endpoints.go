package tclient

import (
	"github.com/spf13/viper"
	"github.com/todarch/todarch-cli/consts"
	"strings"
)

func TodarchApiBase() string {
	api := viper.GetString(consts.TodarchApiBase)
	if strings.HasSuffix(api, "/") {
		return api
	} else {
		return api + "/"
	}
}

const (
	tdPrefix            = "td/"
	tdUp                = tdPrefix + "non-secured/up"
	currentUserTodosURL = tdPrefix + "api/todos"
	newTodoURL          = tdPrefix + "api/todos"
	getSingleTodo       = tdPrefix + "api/todos"
)

const (
	umPrefix = "um/"
	umUp     = umPrefix + "non-secured/up"
	loginURL = umPrefix + "non-secured/authenticate"
)
