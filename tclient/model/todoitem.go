package model

type TodoItem struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"description"`
	Priority int    `json:"priority"`
}
