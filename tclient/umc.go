package tclient

import (
	"fmt"
)

type loginReq struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func Authenticate(uname string, pword string) {
	_, err := doReq(requestOptions{
		Method: "POST",
		URL:    loginURL,
		Body:   loginReq{uname, pword},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfuly logged in.")
	}
}
