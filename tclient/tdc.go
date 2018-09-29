package tclient

import (
	"encoding/json"
	"github.com/todarch/todarch-cli/util"
)

func IsTdUp() bool {
	_, err := doReq(requestOptions{URL: tdUp})
	if err != nil {
		return false
	}
	return true
}

func CurrentUserTodos() []TodoItem {
	res, err := doReq(requestOptions{URL: currentUserTodosURL})
	if err != nil {
		util.SayLastWords(err.Error())
	}
	var todos []TodoItem
	err = json.Unmarshal([]byte(res), &todos)
	util.PanicOnError(err)
	return todos
}

func NewTodo(todoReq TodoCreationReq) {
	_, err := doReq(requestOptions{URL: newTodoURL, Method: "POST", Body: todoReq})
	util.PanicOnError(err)
}

func IsUmUp() bool {
	_, err := doReq(requestOptions{URL: umUp})
	if err != nil {
		return false
	}
	return true
}
