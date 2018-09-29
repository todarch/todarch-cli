package tclient

const todarchApiBase = "https://todarch-gw.herokuapp.com/"

const (
	tdPrefix            = todarchApiBase + "td/"
	tdUp                = tdPrefix + "non-secured/up"
	currentUserTodosURL = tdPrefix + "api/todos"
	newTodoURL          = tdPrefix + "api/todos"
)

const (
	umPrefix = todarchApiBase + "um/"
	umUp     = umPrefix + "non-secured/up"
	loginURL = umPrefix + "non-secured/authenticate"
)
