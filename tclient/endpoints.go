package tclient

const todarchApiBase = "https://todarch-gw.herokuapp.com/"

const (
	tdPrefix = todarchApiBase + "td/"
	tdUp     = tdPrefix + "non-secured/up"
)

const (
	umPrefix = todarchApiBase + "um/"
	umUp     = umPrefix + "non-secured/up"
	loginURL = umPrefix + "non-secured/authenticate"
)