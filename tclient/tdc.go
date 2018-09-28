package tclient

import (
	"encoding/json"
	"fmt"
	"github.com/todarch/todarch-cli/tclient/model"
	"github.com/todarch/todarch-cli/util"
)

func IsTdUp() bool {
	_, err := doReq(requestOptions{URL: tdUp})
	if err != nil {
		return false
	}
	return true
}

func CurrentUserTodos() {
	res, err := doReq(requestOptions{URL: currentUserTodosURL})
	if err != nil {
		fmt.Println(err)
		return
	}
	var todos []model.TodoItem
	err = json.Unmarshal([]byte(res), &todos)
	util.PanicOnError(err)
	if len(todos) == 0 {
		fmt.Println("You do not have any todos yet.")
	} else {
		fmt.Println("ID   TITLE     PRIORITY")
		for _, todo := range todos {
			fmt.Println(todo)
		}
	}
}

func IsUmUp() bool {
	_, err := doReq(requestOptions{URL: umUp})
	if err != nil {
		return false
	}
	return true
}
