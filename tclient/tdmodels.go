package tclient

type TodoItem struct {
	Id              int      `json:"id"`
	Title           string   `json:"title"`
	Desc            string   `json:"description"`
	Priority        int      `json:"priority"`
	Status          string   `json:"status"`
	TimeNeededInMin int      `json:"timeNeededInMin"`
	Tags            []string `json:"tags"`
}

type TodoCreationReq struct {
	Title           string   `yaml:"title" json:"title"`
	Description     string   `yaml:"description" json:"description"`
	Priority        int      `yaml:"priority" json:"priority"`
	TimeNeededInMin int      `yaml:"timeNeededInMin" json:"timeNeededInMin"`
	Tags            []string `yaml:"tags" json:"tags"`
}
