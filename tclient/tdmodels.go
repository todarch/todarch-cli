package tclient

type TodoItem struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"description"`
	Priority int    `json:"priority"`
}

type TodoCreationReq struct {
	Title       string `yaml:"title" json:"title"`
	Description string `yaml:"description" json:"description"`
	Priority    int    `yaml:"priority" json:"priority"`
	//Tags        []string `yaml:"tags"`
}
