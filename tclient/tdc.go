package tclient

import (
	"encoding/json"
	"github.com/todarch/todarch-cli/util"
	"strconv"
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

func TodoById(todoId int) TodoItem {
	todoByIdUrl := getSingleTodo + "/" + strconv.Itoa(todoId)
	res, err := doReq(requestOptions{URL: todoByIdUrl})
	if err != nil {
		util.SayLastWords(err.Error())
	}
	var singleTodo TodoItem
	err = json.Unmarshal([]byte(res), &singleTodo)
	util.PanicOnError(err)
	return singleTodo
}

func GetTodoDone(todoId int) {
	getDoneUrl := getSingleTodo + "/" + strconv.Itoa(todoId) + "/done"
	_, err := doReq(requestOptions{URL: getDoneUrl, Method: "PUT"})
	if err != nil {
		util.SayLastWords(err.Error())
	}
	util.Debug("Todo get done successfully!")
}

func DeleteTodoById(todoId int) {
	deleteTodoUrl := getSingleTodo + "/" + strconv.Itoa(todoId)
	_, err := doReq(requestOptions{URL: deleteTodoUrl, Method: "DELETE"})
	if err != nil {
		util.SayLastWords(err.Error())
	}
	util.Debug("Todo removed successfully")
}

func IsUmUp() bool {
	_, err := doReq(requestOptions{URL: umUp})
	if err != nil {
		return false
	}
	return true
}
